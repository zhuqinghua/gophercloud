package tenants

import "gerrit.mcp.mirantis.net/debian/gophercloud.git"

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("tenants")
}
