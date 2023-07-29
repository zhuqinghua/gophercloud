package roles

import "github.com/zhuqinghua/gophercloud/pagination"

// RoleAssignment is the result of a role assignments query.
type RoleAssignment struct {
	Role  Role  `json:"role,omitempty"`
	Scope Scope `json:"scope,omitempty"`
	User  User  `json:"user,omitempty"`
	Group Group `json:"group,omitempty"`
}

type Role struct {
	ID string `json:"id,omitempty"`
}

type Scope struct {
	Domain  Domain  `json:"domain,omitempty"`
	Project Project `json:"project,omitempty"`
}

type Domain struct {
	ID string `json:"id,omitempty"`
}

type Project struct {
	ID string `json:"id,omitempty"`
}

type User struct {
	ID string `json:"id,omitempty"`
}

type Group struct {
	ID string `json:"id,omitempty"`
}

// RoleAssignmentPage is a single page of RoleAssignments results.
type RoleAssignmentPage struct {
	pagination.LinkedPageBase
}

type RolePage struct {
	pagination.SinglePageBase
}

// IsEmpty returns true if the page contains no results.
func (r RoleAssignmentPage) IsEmpty() (bool, error) {
	roleAssignments, err := ExtractRoleAssignments(r)
	return len(roleAssignments) == 0, err
}

// NextPageURL uses the response's embedded link reference to navigate to the next page of results.
func (r RoleAssignmentPage) NextPageURL() (string, error) {
	var s struct {
		Links struct {
			Next string `json:"next"`
		} `json:"links"`
	}
	err := r.ExtractInto(&s)
	return s.Links.Next, err
}

// ExtractRoleAssignments extracts a slice of RoleAssignments from a Collection acquired from List.
func ExtractRoleAssignments(r pagination.Page) ([]RoleAssignment, error) {
	var s struct {
		RoleAssignments []RoleAssignment `json:"role_assignments"`
	}
	err := (r.(RoleAssignmentPage)).ExtractInto(&s)
	return s.RoleAssignments, err
}
func ExtractRoles(r pagination.Page) ([]Role, error) {
	var s struct {
		Roles []Role `json:"roles"`
	}
	err := (r.(RolePage)).ExtractInto(&s)
	return s.Roles, err
}
