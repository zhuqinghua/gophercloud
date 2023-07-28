package accounts

import "gerrit.mcp.mirantis.net/debian/gophercloud.git"

func getURL(c *gophercloud.ServiceClient) string {
	return c.Endpoint
}

func updateURL(c *gophercloud.ServiceClient) string {
	return getURL(c)
}
