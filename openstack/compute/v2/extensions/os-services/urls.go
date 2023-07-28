package os_services

import gophercloud "github.com/zhuqinghua/gophercloud.git"

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-services")
}
