package os_services

import gophercloud "gerrit.mcp.mirantis.net/debian/gophercloud.git"

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-services")
}
