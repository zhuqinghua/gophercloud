package testing

import (
	"testing"

	"gerrit.mcp.mirantis.net/debian/gophercloud.git"
	th "gerrit.mcp.mirantis.net/debian/gophercloud.git/testhelper"
)

func TestServiceURL(t *testing.T) {
	c := &gophercloud.ServiceClient{Endpoint: "http://123.45.67.8/"}
	expected := "http://123.45.67.8/more/parts/here"
	actual := c.ServiceURL("more", "parts", "here")
	th.CheckEquals(t, expected, actual)
}
