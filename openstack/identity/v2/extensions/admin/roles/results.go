package roles

import (
	"gerrit.mcp.mirantis.net/debian/gophercloud.git"
	"gerrit.mcp.mirantis.net/debian/gophercloud.git/pagination"
)

// Role represents an API role resource.
type Role struct {
	// The unique ID for the role.
	ID string

	// The human-readable name of the role.
	Name string

	// The description of the role.
	Description string

	// The associated service for this role.
	ServiceID string
}

// RolePage is a single page of a user Role collection.
type RolePage struct {
	pagination.SinglePageBase
}

// IsEmpty determines whether or not a page of Tenants contains any results.
func (r RolePage) IsEmpty() (bool, error) {
	users, err := ExtractRoles(r)
	return len(users) == 0, err
}

// ExtractRoles returns a slice of roles contained in a single page of results.
func ExtractRoles(r pagination.Page) ([]Role, error) {
	var s struct {
		Roles []Role `json:"roles"`
	}
	err := (r.(RolePage)).ExtractInto(&s)
	return s.Roles, err
}

// UserRoleResult represents the result of either an AddUserRole or
// a DeleteUserRole operation.
type UserRoleResult struct {
	gophercloud.ErrResult
}
