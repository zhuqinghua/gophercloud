package buildinfo

import "gerrit.mcp.mirantis.net/debian/gophercloud.git"

func getURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("build_info")
}
