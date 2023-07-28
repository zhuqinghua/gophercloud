package testing

import (
	"testing"

	common "gerrit.mcp.mirantis.net/debian/gophercloud.git/openstack/common/extensions/testing"
	"gerrit.mcp.mirantis.net/debian/gophercloud.git/openstack/identity/v2/extensions"
	"gerrit.mcp.mirantis.net/debian/gophercloud.git/pagination"
	th "gerrit.mcp.mirantis.net/debian/gophercloud.git/testhelper"
	"gerrit.mcp.mirantis.net/debian/gophercloud.git/testhelper/client"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListExtensionsSuccessfully(t)

	count := 0
	err := extensions.List(client.ServiceClient()).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := extensions.ExtractExtensions(page)
		th.AssertNoErr(t, err)
		th.CheckDeepEquals(t, common.ExpectedExtensions, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, count)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	common.HandleGetExtensionSuccessfully(t)

	actual, err := extensions.Get(client.ServiceClient(), "agent").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, common.SingleExtension, actual)
}
