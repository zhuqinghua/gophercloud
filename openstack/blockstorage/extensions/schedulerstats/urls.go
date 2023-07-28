package schedulerstats

func storagePoolsListURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("scheduler-stats", "get_pools")
}
