package zones

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("zones")
}

func zoneURL(c *gophercloud.ServiceClient, zoneID string) string {
	return c.ServiceURL("zones", zoneID)
}
