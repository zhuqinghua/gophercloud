// +build acceptance

package v1

import (
	"testing"

	"gerrit.mcp.mirantis.net/debian/gophercloud.git"
	"gerrit.mcp.mirantis.net/debian/gophercloud.git/openstack/orchestration/v1/stackevents"
	"gerrit.mcp.mirantis.net/debian/gophercloud.git/openstack/orchestration/v1/stacks"
	"gerrit.mcp.mirantis.net/debian/gophercloud.git/pagination"
	th "gerrit.mcp.mirantis.net/debian/gophercloud.git/testhelper"
)

func TestStackEvents(t *testing.T) {
	// Create a provider client for making the HTTP requests.
	// See common.go in this directory for more information.
	client := newClient(t)

	stackName := "postman_stack_2"
	resourceName := "hello_world"
	var eventID string

	createOpts := stacks.CreateOpts{
		Name:     stackName,
		Template: template,
		Timeout:  5,
	}
	stack, err := stacks.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Created stack: %+v\n", stack)
	defer func() {
		err := stacks.Delete(client, stackName, stack.ID).ExtractErr()
		th.AssertNoErr(t, err)
		t.Logf("Deleted stack (%s)", stackName)
	}()
	err = gophercloud.WaitFor(60, func() (bool, error) {
		getStack, err := stacks.Get(client, stackName, stack.ID).Extract()
		if err != nil {
			return false, err
		}
		if getStack.Status == "CREATE_COMPLETE" {
			return true, nil
		}
		return false, nil
	})

	err = stackevents.List(client, stackName, stack.ID, nil).EachPage(func(page pagination.Page) (bool, error) {
		events, err := stackevents.ExtractEvents(page)
		th.AssertNoErr(t, err)
		t.Logf("listed events: %+v\n", events)
		eventID = events[0].ID
		return false, nil
	})
	th.AssertNoErr(t, err)

	err = stackevents.ListResourceEvents(client, stackName, stack.ID, resourceName, nil).EachPage(func(page pagination.Page) (bool, error) {
		resourceEvents, err := stackevents.ExtractEvents(page)
		th.AssertNoErr(t, err)
		t.Logf("listed resource events: %+v\n", resourceEvents)
		return false, nil
	})
	th.AssertNoErr(t, err)

	event, err := stackevents.Get(client, stackName, stack.ID, resourceName, eventID).Extract()
	th.AssertNoErr(t, err)
	t.Logf("retrieved event: %+v\n", event)
}