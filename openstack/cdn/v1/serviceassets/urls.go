package serviceassets

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("services", id, "assets")
}
