package snapshots

import (
	gophercloud "github.com/zhuqinghua/gophercloud.git"
	"github.com/zhuqinghua/gophercloud.git/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the List
// request.
type ListOptsBuilder interface {
	ToVolumeListQuery() (string, error)
}

// ListOpts holds options for listing Snapshots. It is passed to the snapshot.List
// function.
type ListOpts struct {
	// admin-only option. Set it to true to see all tenant volumes.
	AllTenants bool `q:"all_tenants"`
}

// ToVolumeListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToVolumeListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns Volumes optionally limited by the conditions provided in ListOpts.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToVolumeListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return SnapshotPage{pagination.SinglePageBase(r)}
	})
}
