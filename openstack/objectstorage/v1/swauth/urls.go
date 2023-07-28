package swauth

import "gerrit.mcp.mirantis.net/debian/gophercloud.git"

func getURL(c *gophercloud.ProviderClient) string {
	return c.IdentityBase + "auth/v1.0"
}
