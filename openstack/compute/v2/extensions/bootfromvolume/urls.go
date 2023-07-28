package bootfromvolume

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-volumes_boot")
}
