package agents

import gophercloud "github.com/zhuqinghua/gophercloud.git"

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("agents")
}

func listDHCPNetworksURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("agents", id, "dhcp-networks")
}
