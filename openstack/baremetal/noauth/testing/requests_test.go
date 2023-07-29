package testing

import (
	"testing"

	"github.com/zhuqinghua/gophercloud/openstack/baremetal/noauth"
	th "github.com/zhuqinghua/gophercloud/testhelper"
)

func TestNoAuth(t *testing.T) {
	noauthClient, err := noauth.NewBareMetalNoAuth(noauth.EndpointOpts{
		IronicEndpoint: "http://ironic:6385/v1",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "", noauthClient.TokenID)
}
