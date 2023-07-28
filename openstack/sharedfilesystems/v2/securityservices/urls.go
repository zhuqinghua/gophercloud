package securityservices

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("security-services")
}

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("security-services", id)
}

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("security-services", "detail")
}
