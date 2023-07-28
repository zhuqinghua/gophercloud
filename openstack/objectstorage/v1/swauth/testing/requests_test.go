package testing

import (
	"testing"

	"gerrit.mcp.mirantis.net/debian/gophercloud.git/openstack"
	"gerrit.mcp.mirantis.net/debian/gophercloud.git/openstack/objectstorage/v1/swauth"
	th "gerrit.mcp.mirantis.net/debian/gophercloud.git/testhelper"
)

func TestAuth(t *testing.T) {
	authOpts := swauth.AuthOpts{
		User: "test:tester",
		Key:  "testing",
	}

	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleAuthSuccessfully(t, authOpts)

	providerClient, err := openstack.NewClient(th.Endpoint())
	th.AssertNoErr(t, err)

	swiftClient, err := swauth.NewObjectStorageV1(providerClient, authOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, swiftClient.TokenID, AuthResult.Token)
}
