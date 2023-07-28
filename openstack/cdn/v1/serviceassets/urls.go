package serviceassets

import "gerrit.mcp.mirantis.net/debian/gophercloud.git"

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("services", id, "assets")
}
