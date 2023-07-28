package apiversions

import (
	"gerrit.mcp.mirantis.net/debian/gophercloud.git"
)

func getURL(c *gophercloud.ServiceClient, version string) string {
	return c.ServiceURL(version)
}

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL()
}
