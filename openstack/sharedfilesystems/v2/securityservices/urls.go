package securityservices

import "gerrit.mcp.mirantis.net/debian/gophercloud.git"

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("security-services")
}

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("security-services", id)
}

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("security-services", "detail")
}
