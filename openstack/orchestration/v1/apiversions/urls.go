package apiversions

func apiVersionsURL(c *gophercloud.ServiceClient) string {
	return c.Endpoint
}
