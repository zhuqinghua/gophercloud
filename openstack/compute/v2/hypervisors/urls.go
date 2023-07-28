package hypervisors

func hypervisorsListDetailURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-hypervisors", "detail")
}

func aggregatesListURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-aggregates")
}
