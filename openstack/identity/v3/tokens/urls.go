package tokens

func tokenURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("auth", "tokens")
}
