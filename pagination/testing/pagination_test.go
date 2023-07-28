package testing

import (
	"github.com/zhuqinghua/gophercloud.git/testhelper"
)

func createClient() *gophercloud.ServiceClient {
	return &gophercloud.ServiceClient{
		ProviderClient: &gophercloud.ProviderClient{TokenID: "abc123"},
		Endpoint:       testhelper.Endpoint(),
	}
}
