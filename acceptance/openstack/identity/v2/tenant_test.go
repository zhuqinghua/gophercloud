//go:build acceptance || identity
// +build acceptance identity

package v2

import (
	"testing"

	"github.com/zhuqinghua/gophercloud/acceptance/clients"
	"github.com/zhuqinghua/gophercloud/acceptance/tools"
	"github.com/zhuqinghua/gophercloud/openstack/identity/v2/tenants"
)

func TestTenantsList(t *testing.T) {
	client, err := clients.NewIdentityV2Client()
	if err != nil {
		t.Fatalf("Unable to obtain an identity client: %v")
	}

	allPages, err := tenants.List(client, nil).AllPages()
	if err != nil {
		t.Fatalf("Unable to list tenants: %v", err)
	}

	allTenants, err := tenants.ExtractTenants(allPages)
	if err != nil {
		t.Fatalf("Unable to extract tenants: %v", err)
	}

	for _, tenant := range allTenants {
		tools.PrintResource(t, tenant)
	}
}
