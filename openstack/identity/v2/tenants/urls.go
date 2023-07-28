package tenants

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("tenants")
}
