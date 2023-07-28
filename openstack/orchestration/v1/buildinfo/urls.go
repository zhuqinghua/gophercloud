package buildinfo

func getURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("build_info")
}
