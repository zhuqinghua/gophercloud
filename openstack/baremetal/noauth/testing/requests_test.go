package testing

import (
	"testing"

	"gerrit.mcp.mirantis.net/debian/gophercloud.git/openstack/baremetal/noauth"
	th "gerrit.mcp.mirantis.net/debian/gophercloud.git/testhelper"
)

func TestNoAuth(t *testing.T) {
	noauthClient, err := noauth.NewBareMetalNoAuth(noauth.EndpointOpts{
		IronicEndpoint: "http://ironic:6385/v1",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "", noauthClient.TokenID)
}
