package limits

const resourcePath = "limits"

func getURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}
