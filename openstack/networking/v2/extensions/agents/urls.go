package agents

import gophercloud "gerrit.mcp.mirantis.net/debian/gophercloud.git"

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("agents")
}

func listDHCPNetworksURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("agents", id, "dhcp-networks")
}
