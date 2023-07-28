package extensions

import (
	"gerrit.mcp.mirantis.net/debian/gophercloud.git"
	common "gerrit.mcp.mirantis.net/debian/gophercloud.git/openstack/common/extensions"
	"gerrit.mcp.mirantis.net/debian/gophercloud.git/pagination"
)

// ExtractExtensions interprets a Page as a slice of Extensions.
func ExtractExtensions(page pagination.Page) ([]common.Extension, error) {
	return common.ExtractExtensions(page)
}

// Get retrieves information for a specific extension using its alias.
func Get(c *gophercloud.ServiceClient, alias string) common.GetResult {
	return common.Get(c, alias)
}

// List returns a Pager which allows you to iterate over the full collection of extensions.
// It does not accept query parameters.
func List(c *gophercloud.ServiceClient) pagination.Pager {
	return common.List(c)
}
