//go:build acceptance || networking
// +build acceptance networking

package v2

import (
	"testing"

	"github.com/zhuqinghua/gophercloud/acceptance/clients"
	extensions "github.com/zhuqinghua/gophercloud/acceptance/openstack/networking/v2/extensions"
	"github.com/zhuqinghua/gophercloud/acceptance/tools"
	"github.com/zhuqinghua/gophercloud/openstack/networking/v2/ports"
)

func TestPortsList(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	if err != nil {
		t.Fatalf("Unable to create a network client: %v", err)
	}

	allPages, err := ports.List(client, nil).AllPages()
	if err != nil {
		t.Fatalf("Unable to list ports: %v", err)
	}

	allPorts, err := ports.ExtractPorts(allPages)
	if err != nil {
		t.Fatalf("Unable to extract ports: %v", err)
	}

	for _, port := range allPorts {
		tools.PrintResource(t, port)
	}
}

func TestPortsCRUD(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	if err != nil {
		t.Fatalf("Unable to create a network client: %v", err)
	}

	// Create Network
	network, err := CreateNetwork(t, client)
	if err != nil {
		t.Fatalf("Unable to create network: %v", err)
	}
	defer DeleteNetwork(t, client, network.ID)

	// Create Subnet
	subnet, err := CreateSubnet(t, client, network.ID)
	if err != nil {
		t.Fatalf("Unable to create subnet: %v", err)
	}
	defer DeleteSubnet(t, client, subnet.ID)

	// Create port
	port, err := CreatePort(t, client, network.ID, subnet.ID)
	if err != nil {
		t.Fatalf("Unable to create port: %v", err)
	}
	defer DeletePort(t, client, port.ID)

	tools.PrintResource(t, port)

	// Update port
	newPortName := tools.RandomString("TESTACC-", 8)
	updateOpts := ports.UpdateOpts{
		Name: newPortName,
	}
	newPort, err := ports.Update(client, port.ID, updateOpts).Extract()
	if err != nil {
		t.Fatalf("Could not update port: %v", err)
	}

	tools.PrintResource(t, newPort)
}

func TestPortsRemoveSecurityGroups(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	if err != nil {
		t.Fatalf("Unable to create a network client: %v", err)
	}

	// Create Network
	network, err := CreateNetwork(t, client)
	if err != nil {
		t.Fatalf("Unable to create network: %v", err)
	}
	defer DeleteNetwork(t, client, network.ID)

	// Create Subnet
	subnet, err := CreateSubnet(t, client, network.ID)
	if err != nil {
		t.Fatalf("Unable to create subnet: %v", err)
	}
	defer DeleteSubnet(t, client, subnet.ID)

	// Create port
	port, err := CreatePort(t, client, network.ID, subnet.ID)
	if err != nil {
		t.Fatalf("Unable to create port: %v", err)
	}
	defer DeletePort(t, client, port.ID)

	tools.PrintResource(t, port)

	// Create a Security Group
	group, err := extensions.CreateSecurityGroup(t, client)
	if err != nil {
		t.Fatalf("Unable to create security group: %v", err)
	}
	defer extensions.DeleteSecurityGroup(t, client, group.ID)

	// Add the group to the port
	updateOpts := ports.UpdateOpts{
		SecurityGroups: []string{group.ID},
	}
	newPort, err := ports.Update(client, port.ID, updateOpts).Extract()
	if err != nil {
		t.Fatalf("Could not update port: %v", err)
	}

	// Remove the group
	updateOpts = ports.UpdateOpts{
		SecurityGroups: []string{},
	}
	newPort, err = ports.Update(client, port.ID, updateOpts).Extract()
	if err != nil {
		t.Fatalf("Could not update port: %v", err)
	}

	tools.PrintResource(t, newPort)

	if len(newPort.SecurityGroups) > 0 {
		t.Fatalf("Unable to remove security group from port")
	}
}

func TestPortsRemoveAddressPair(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	if err != nil {
		t.Fatalf("Unable to create a network client: %v", err)
	}

	// Create Network
	network, err := CreateNetwork(t, client)
	if err != nil {
		t.Fatalf("Unable to create network: %v", err)
	}
	defer DeleteNetwork(t, client, network.ID)

	// Create Subnet
	subnet, err := CreateSubnet(t, client, network.ID)
	if err != nil {
		t.Fatalf("Unable to create subnet: %v", err)
	}
	defer DeleteSubnet(t, client, subnet.ID)

	// Create port
	port, err := CreatePort(t, client, network.ID, subnet.ID)
	if err != nil {
		t.Fatalf("Unable to create port: %v", err)
	}
	defer DeletePort(t, client, port.ID)

	tools.PrintResource(t, port)

	// Add an address pair to the port
	updateOpts := ports.UpdateOpts{
		AllowedAddressPairs: []ports.AddressPair{
			ports.AddressPair{IPAddress: "192.168.255.10", MACAddress: "aa:bb:cc:dd:ee:ff"},
		},
	}
	newPort, err := ports.Update(client, port.ID, updateOpts).Extract()
	if err != nil {
		t.Fatalf("Could not update port: %v", err)
	}

	// Remove the address pair
	updateOpts = ports.UpdateOpts{
		AllowedAddressPairs: []ports.AddressPair{},
	}
	newPort, err = ports.Update(client, port.ID, updateOpts).Extract()
	if err != nil {
		t.Fatalf("Could not update port: %v", err)
	}

	tools.PrintResource(t, newPort)

	if len(newPort.AllowedAddressPairs) > 0 {
		t.Fatalf("Unable to remove the address pair")
	}
}
