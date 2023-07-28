package testing

import (
	"testing"

	"gerrit.mcp.mirantis.net/debian/gophercloud.git/openstack/orchestration/v1/buildinfo"
	th "gerrit.mcp.mirantis.net/debian/gophercloud.git/testhelper"
	fake "gerrit.mcp.mirantis.net/debian/gophercloud.git/testhelper/client"
)

func TestGetTemplate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t, GetOutput)

	actual, err := buildinfo.Get(fake.ServiceClient()).Extract()
	th.AssertNoErr(t, err)

	expected := GetExpected
	th.AssertDeepEquals(t, expected, actual)
}
