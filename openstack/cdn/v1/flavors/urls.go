package flavors

import "gerrit.mcp.mirantis.net/debian/gophercloud.git"

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("flavors")
}

func getURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("flavors", id)
}