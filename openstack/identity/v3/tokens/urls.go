package tokens

import "gerrit.mcp.mirantis.net/debian/gophercloud.git"

func tokenURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("auth", "tokens")
}
