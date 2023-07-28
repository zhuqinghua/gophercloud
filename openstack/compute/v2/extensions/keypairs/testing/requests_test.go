package testing

import (
	"testing"

	"gerrit.mcp.mirantis.net/debian/gophercloud.git/openstack/compute/v2/extensions/keypairs"
	"gerrit.mcp.mirantis.net/debian/gophercloud.git/pagination"
	th "gerrit.mcp.mirantis.net/debian/gophercloud.git/testhelper"
	"gerrit.mcp.mirantis.net/debian/gophercloud.git/testhelper/client"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	count := 0
	err := keypairs.List(client.ServiceClient()).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := keypairs.ExtractKeyPairs(page)
		th.AssertNoErr(t, err)
		th.CheckDeepEquals(t, ExpectedKeyPairSlice, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, count)
}

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	actual, err := keypairs.Create(client.ServiceClient(), keypairs.CreateOpts{
		Name: "createdkey",
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreatedKeyPair, actual)
}

func TestImport(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleImportSuccessfully(t)

	actual, err := keypairs.Create(client.ServiceClient(), keypairs.CreateOpts{
		Name:      "importedkey",
		PublicKey: "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQDx8nkQv/zgGgB4rMYmIf+6A4l6Rr+o/6lHBQdW5aYd44bd8JttDCE/F/pNRr0lRE+PiqSPO8nDPHw0010JeMH9gYgnnFlyY3/OcJ02RhIPyyxYpv9FhY+2YiUkpwFOcLImyrxEsYXpD/0d3ac30bNH6Sw9JD9UZHYcpSxsIbECHw== Generated by Nova",
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ImportedKeyPair, actual)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	actual, err := keypairs.Get(client.ServiceClient(), "firstkey").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &FirstKeyPair, actual)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteSuccessfully(t)

	err := keypairs.Delete(client.ServiceClient(), "deletedkey").ExtractErr()
	th.AssertNoErr(t, err)
}