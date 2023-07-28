package members

import "gerrit.mcp.mirantis.net/debian/gophercloud.git"

const (
	rootPath     = "lb"
	resourcePath = "members"
)

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(rootPath, resourcePath)
}

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id)
}
