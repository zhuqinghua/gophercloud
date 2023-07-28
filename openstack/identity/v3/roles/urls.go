package roles

import "gerrit.mcp.mirantis.net/debian/gophercloud.git"

func listAssignmentsURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("role_assignments")
}

func listURL(client *gophercloud.ServiceClient) string{
    return client.ServiceURL("roles")
}
