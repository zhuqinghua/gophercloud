package flavors

func getURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("flavors", id)
}

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("flavors", "detail")
}

func createURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("flavors")
}
