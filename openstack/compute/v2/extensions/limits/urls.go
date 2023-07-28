package limits

import (
	"gerrit.mcp.mirantis.net/debian/gophercloud.git"
)

const resourcePath = "limits"

func getURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}
