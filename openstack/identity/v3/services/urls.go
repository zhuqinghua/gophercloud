package services

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("services")
}

func serviceURL(client *gophercloud.ServiceClient, serviceID string) string {
	return client.ServiceURL("services", serviceID)
}
