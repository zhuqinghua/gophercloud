package stackresources

import "gerrit.mcp.mirantis.net/debian/gophercloud.git"

func findURL(c *gophercloud.ServiceClient, stackName string) string {
	return c.ServiceURL("stacks", stackName, "resources")
}

func listURL(c *gophercloud.ServiceClient, stackName, stackID string) string {
	return c.ServiceURL("stacks", stackName, stackID, "resources")
}

func getURL(c *gophercloud.ServiceClient, stackName, stackID, resourceName string) string {
	return c.ServiceURL("stacks", stackName, stackID, "resources", resourceName)
}

func metadataURL(c *gophercloud.ServiceClient, stackName, stackID, resourceName string) string {
	return c.ServiceURL("stacks", stackName, stackID, "resources", resourceName, "metadata")
}

func listTypesURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("resource_types")
}

func schemaURL(c *gophercloud.ServiceClient, typeName string) string {
	return c.ServiceURL("resource_types", typeName)
}

func templateURL(c *gophercloud.ServiceClient, typeName string) string {
	return c.ServiceURL("resource_types", typeName, "template")
}
