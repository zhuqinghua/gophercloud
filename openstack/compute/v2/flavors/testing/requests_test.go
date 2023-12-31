package testing

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/zhuqinghua/gophercloud/openstack/compute/v2/flavors"
	"github.com/zhuqinghua/gophercloud/pagination"
	th "github.com/zhuqinghua/gophercloud/testhelper"
	fake "github.com/zhuqinghua/gophercloud/testhelper/client"
)

const tokenID = "blerb"

func TestListFlavors(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/flavors/detail", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		r.ParseForm()
		marker := r.Form.Get("marker")
		switch marker {
		case "":
			fmt.Fprintf(w, `
					{
						"flavors": [
							{
								"id": "1",
								"name": "m1.tiny",
								"disk": 1,
								"ram": 512,
								"vcpus": 1,
								"swap":""
							},
							{
								"id": "2",
								"name": "m2.small",
								"disk": 10,
								"ram": 1024,
								"vcpus": 2,
								"swap": 1000
							}
						],
						"flavors_links": [
							{
								"href": "%s/flavors/detail?marker=2",
								"rel": "next"
							}
						]
					}
				`, th.Server.URL)
		case "2":
			fmt.Fprintf(w, `{ "flavors": [] }`)
		default:
			t.Fatalf("Unexpected marker: [%s]", marker)
		}
	})

	pages := 0
	err := flavors.ListDetail(fake.ServiceClient(), nil).EachPage(func(page pagination.Page) (bool, error) {
		pages++

		actual, err := flavors.ExtractFlavors(page)
		if err != nil {
			return false, err
		}

		expected := []flavors.Flavor{
			{ID: "1", Name: "m1.tiny", Disk: 1, RAM: 512, VCPUs: 1, Swap: 0},
			{ID: "2", Name: "m2.small", Disk: 10, RAM: 1024, VCPUs: 2, Swap: 1000},
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Expected %#v, but was %#v", expected, actual)
		}

		return true, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if pages != 1 {
		t.Errorf("Expected one page, got %d", pages)
	}
}

func TestGetFlavor(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/flavors/12345", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, `
			{
				"flavor": {
					"id": "1",
					"name": "m1.tiny",
					"disk": 1,
					"ram": 512,
					"vcpus": 1,
					"rxtx_factor": 1,
					"swap": ""
				}
			}
		`)
	})

	actual, err := flavors.Get(fake.ServiceClient(), "12345").Extract()
	if err != nil {
		t.Fatalf("Unable to get flavor: %v", err)
	}

	expected := &flavors.Flavor{
		ID:         "1",
		Name:       "m1.tiny",
		Disk:       1,
		RAM:        512,
		VCPUs:      1,
		RxTxFactor: 1,
		Swap:       0,
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %#v, but was %#v", expected, actual)
	}
}

func TestCreateFlavor(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/flavors", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, `
			{
				"flavor": {
					"id": "1",
					"name": "m1.tiny",
					"disk": 1,
					"ram": 512,
					"vcpus": 1,
					"rxtx_factor": 1,
					"swap": ""
				}
			}
		`)
	})

	disk := 1
	opts := &flavors.CreateOpts{
		ID:         "1",
		Name:       "m1.tiny",
		Disk:       &disk,
		RAM:        512,
		VCPUs:      1,
		RxTxFactor: 1.0,
	}
	actual, err := flavors.Create(fake.ServiceClient(), opts).Extract()
	if err != nil {
		t.Fatalf("Unable to create flavor: %v", err)
	}

	expected := &flavors.Flavor{
		ID:         "1",
		Name:       "m1.tiny",
		Disk:       1,
		RAM:        512,
		VCPUs:      1,
		RxTxFactor: 1,
		Swap:       0,
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %#v, but was %#v", expected, actual)
	}
}
