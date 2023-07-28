package v2

import (
	"github.com/zhuqinghua/gophercloud.git/acceptance/clients"
	"github.com/zhuqinghua/gophercloud.git/openstack/sharedfilesystems/v2/shares"
	"testing"
)

func TestShareCreate(t *testing.T) {
	client, err := clients.NewSharedFileSystemV2Client()
	if err != nil {
		t.Fatalf("Unable to create a sharedfs client: %v", err)
	}

	share, err := CreateShare(t, client)
	if err != nil {
		t.Fatalf("Unable to create a share: %v", err)
	}

	defer DeleteShare(t, client, share)

	created, err := shares.Get(client, share.ID).Extract()
	if err != nil {
		t.Errorf("Unable to retrieve share: %v", err)
	}
	PrintShare(t, created)
}
