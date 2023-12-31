//go:build acceptance || compute || images
// +build acceptance compute images

package v2

import (
	"testing"

	"github.com/zhuqinghua/gophercloud/acceptance/clients"
	"github.com/zhuqinghua/gophercloud/acceptance/tools"
	"github.com/zhuqinghua/gophercloud/openstack/compute/v2/images"
)

func TestImagesList(t *testing.T) {
	client, err := clients.NewComputeV2Client()
	if err != nil {
		t.Fatalf("Unable to create a compute: client: %v", err)
	}

	allPages, err := images.ListDetail(client, nil).AllPages()
	if err != nil {
		t.Fatalf("Unable to retrieve images: %v", err)
	}

	allImages, err := images.ExtractImages(allPages)
	if err != nil {
		t.Fatalf("Unable to extract image results: %v", err)
	}

	for _, image := range allImages {
		tools.PrintResource(t, image)
	}
}

func TestImagesGet(t *testing.T) {
	client, err := clients.NewComputeV2Client()
	if err != nil {
		t.Fatalf("Unable to create a compute: client: %v", err)
	}

	choices, err := clients.AcceptanceTestChoicesFromEnv()
	if err != nil {
		t.Fatal(err)
	}

	image, err := images.Get(client, choices.ImageID).Extract()
	if err != nil {
		t.Fatalf("Unable to get image information: %v", err)
	}

	tools.PrintResource(t, image)
}
