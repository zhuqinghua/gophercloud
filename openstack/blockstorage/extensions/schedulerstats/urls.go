package schedulerstats

import "gerrit.mcp.mirantis.net/debian/gophercloud.git"

func storagePoolsListURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("scheduler-stats", "get_pools")
}
