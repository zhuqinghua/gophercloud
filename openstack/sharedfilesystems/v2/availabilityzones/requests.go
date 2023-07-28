package availabilityzones

import (
	"gerrit.mcp.mirantis.net/debian/gophercloud.git"
	"gerrit.mcp.mirantis.net/debian/gophercloud.git/pagination"
)

// List will return the existing availability zones.
func List(client *gophercloud.ServiceClient) pagination.Pager {
	return pagination.NewPager(client, listURL(client), func(r pagination.PageResult) pagination.Page {
		return AvailabilityZonePage{pagination.SinglePageBase(r)}
	})
}
