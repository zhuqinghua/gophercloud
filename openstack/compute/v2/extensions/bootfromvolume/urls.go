package bootfromvolume

import "gerrit.mcp.mirantis.net/debian/gophercloud.git"

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-volumes_boot")
}
