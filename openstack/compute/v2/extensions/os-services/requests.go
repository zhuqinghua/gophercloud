package os_services

import (
	gophercloud "gerrit.mcp.mirantis.net/debian/gophercloud.git"
	"gerrit.mcp.mirantis.net/debian/gophercloud.git/pagination"
)

// List returns a Pager that allows you to iterate over a collection of Network.
func List(client *gophercloud.ServiceClient) pagination.Pager {
	return pagination.NewPager(client, listURL(client), func(r pagination.PageResult) pagination.Page {
		return OsServicePage{pagination.SinglePageBase(r)}
	})
}
