package apiversions

import "gerrit.mcp.mirantis.net/debian/gophercloud.git"

func apiVersionsURL(c *gophercloud.ServiceClient) string {
	return c.Endpoint
}
