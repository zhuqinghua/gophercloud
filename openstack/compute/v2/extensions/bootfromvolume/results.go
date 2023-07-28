package bootfromvolume

import (
	os "gerrit.mcp.mirantis.net/debian/gophercloud.git/openstack/compute/v2/servers"
)

// CreateResult temporarily contains the response from a Create call.
type CreateResult struct {
	os.CreateResult
}
