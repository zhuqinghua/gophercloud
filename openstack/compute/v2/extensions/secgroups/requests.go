package secgroups

import (
	"github.com/zhuqinghua/gophercloud/pagination"
)

func commonList(client *gophercloud.ServiceClient, url string) pagination.Pager {
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return SecurityGroupPage{pagination.SinglePageBase(r)}
	})
}

// List will return a collection of all the security groups for a particular
// tenant.
func List(client *gophercloud.ServiceClient) pagination.Pager {
	return commonList(client, rootURL(client))
}

// ListByServer will return a collection of all the security groups which are
// associated with a particular server.
func ListByServer(client *gophercloud.ServiceClient, serverID string) pagination.Pager {
	return commonList(client, listByServerURL(client, serverID))
}

// GroupOpts is the underlying struct responsible for creating or updating
// security groups. It therefore represents the mutable attributes of a
// security group.
type GroupOpts struct {
	// the name of your security group.
	Name string `json:"name" required:"true"`
	// the description of your security group.
	Description string `json:"description" required:"true"`
}

// CreateOpts is the struct responsible for creating a security group.
type CreateOpts GroupOpts

// CreateOptsBuilder builds the create options into a serializable format.
type CreateOptsBuilder interface {
	ToSecGroupCreateMap() (map[string]interface{}, error)
}

// ToSecGroupCreateMap builds the create options into a serializable format.
func (opts CreateOpts) ToSecGroupCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "security_group")
}

// Create will create a new security group.
func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToSecGroupCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(rootURL(client), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// UpdateOpts is the struct responsible for updating an existing security group.
type UpdateOpts GroupOpts

// UpdateOptsBuilder builds the update options into a serializable format.
type UpdateOptsBuilder interface {
	ToSecGroupUpdateMap() (map[string]interface{}, error)
}

// ToSecGroupUpdateMap builds the update options into a serializable format.
func (opts UpdateOpts) ToSecGroupUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "security_group")
}

// Update will modify the mutable properties of a security group, notably its
// name and description.
func Update(client *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToSecGroupUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(resourceURL(client, id), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Get will return details for a particular security group.
func Get(client *gophercloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(resourceURL(client, id), &r.Body, nil)
	return
}

// Delete will permanently delete a security group from the project.
func Delete(client *gophercloud.ServiceClient, id string) (r gophercloud.ErrResult) {
	_, r.Err = client.Delete(resourceURL(client, id), nil)
	return
}

// CreateRuleOpts represents the configuration for adding a new rule to an
// existing security group.
type CreateRuleOpts struct {
	// the ID of the group that this rule will be added to.
	ParentGroupID string `json:"parent_group_id" required:"true"`
	// the lower bound of the port range that will be opened.
	FromPort int `json:"from_port"`
	// the upper bound of the port range that will be opened.
	ToPort int `json:"to_port"`
	// the protocol type that will be allowed, e.g. TCP.
	IPProtocol string `json:"ip_protocol" required:"true"`
	// ONLY required if FromGroupID is blank. This represents the IP range that
	// will be the source of network traffic to your security group. Use
	// 0.0.0.0/0 to allow all IP addresses.
	CIDR string `json:"cidr,omitempty" or:"FromGroupID"`
	// ONLY required if CIDR is blank. This value represents the ID of a group
	// that forwards traffic to the parent group. So, instead of accepting
	// network traffic from an entire IP range, you can instead refine the
	// inbound source by an existing security group.
	FromGroupID string `json:"group_id,omitempty" or:"CIDR"`
}

// CreateRuleOptsBuilder builds the create rule options into a serializable format.
type CreateRuleOptsBuilder interface {
	ToRuleCreateMap() (map[string]interface{}, error)
}

// ToRuleCreateMap builds the create rule options into a serializable format.
func (opts CreateRuleOpts) ToRuleCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "security_group_rule")
}

// CreateRule will add a new rule to an existing security group (whose ID is
// specified in CreateRuleOpts). You have the option of controlling inbound
// traffic from either an IP range (CIDR) or from another security group.
func CreateRule(client *gophercloud.ServiceClient, opts CreateRuleOptsBuilder) (r CreateRuleResult) {
	b, err := opts.ToRuleCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(rootRuleURL(client), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// DeleteRule will permanently delete a rule from a security group.
func DeleteRule(client *gophercloud.ServiceClient, id string) (r gophercloud.ErrResult) {
	_, r.Err = client.Delete(resourceRuleURL(client, id), nil)
	return
}

func actionMap(prefix, groupName string) map[string]map[string]string {
	return map[string]map[string]string{
		prefix + "SecurityGroup": map[string]string{"name": groupName},
	}
}

// AddServer will associate a server and a security group, enforcing the
// rules of the group on the server.
func AddServer(client *gophercloud.ServiceClient, serverID, groupName string) (r gophercloud.ErrResult) {
	_, r.Err = client.Post(serverActionURL(client, serverID), actionMap("add", groupName), &r.Body, nil)
	return
}

// RemoveServer will disassociate a server from a security group.
func RemoveServer(client *gophercloud.ServiceClient, serverID, groupName string) (r gophercloud.ErrResult) {
	_, r.Err = client.Post(serverActionURL(client, serverID), actionMap("remove", groupName), &r.Body, nil)
	return
}
