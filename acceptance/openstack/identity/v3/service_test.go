//go:build acceptance
// +build acceptance

package v3

import (
	"testing"

	"github.com/zhuqinghua/gophercloud/acceptance/clients"
	"github.com/zhuqinghua/gophercloud/acceptance/tools"
	"github.com/zhuqinghua/gophercloud/openstack/identity/v3/services"
)

func TestServicesList(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	if err != nil {
		t.Fatalf("Unable to obtain an identity client: %v")
	}

	allPages, err := services.List(client, nil).AllPages()
	if err != nil {
		t.Fatalf("Unable to list services: %v", err)
	}

	allServices, err := services.ExtractServices(allPages)
	if err != nil {
		t.Fatalf("Unable to extract services: %v", err)
	}

	for _, service := range allServices {
		tools.PrintResource(t, service)
	}

}
