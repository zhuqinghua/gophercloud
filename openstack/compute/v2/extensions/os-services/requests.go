package os_services

import (
	gophercloud "github.com/zhuqinghua/gophercloud.git"
	"github.com/zhuqinghua/gophercloud.git/pagination"
)

// List returns a Pager that allows you to iterate over a collection of Network.
func List(client *gophercloud.ServiceClient) pagination.Pager {
	return pagination.NewPager(client, listURL(client), func(r pagination.PageResult) pagination.Page {
		return OsServicePage{pagination.SinglePageBase(r)}
	})
}
