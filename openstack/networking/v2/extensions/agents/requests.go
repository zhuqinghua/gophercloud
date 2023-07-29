package agents

import (
	gophercloud "github.com/zhuqinghua/gophercloud"
	"github.com/zhuqinghua/gophercloud/openstack/networking/v2/networks"
	"github.com/zhuqinghua/gophercloud/pagination"
)

func List(c *gophercloud.ServiceClient) pagination.Pager {
	u := listURL(c)
	return pagination.NewPager(c, u, func(r pagination.PageResult) pagination.Page {
		return AgentPage{pagination.SinglePageBase(r)}
	})
}

// ListDHCPNetworks makes a request against the API to list networks that a DHCP agent hosts.
func ListDHCPNetworks(c *gophercloud.ServiceClient, id string) pagination.Pager {
	return pagination.NewPager(c, listDHCPNetworksURL(c, id), func(r pagination.PageResult) pagination.Page {
		return networks.NetworkPage{pagination.LinkedPageBase{PageResult: r}}
	})
}
