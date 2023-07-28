package os_services

import "gerrit.mcp.mirantis.net/debian/gophercloud.git/pagination"

type OsService struct {
	ID     string `json:"id"`
	Host   string `json:"host"`
	State  string `json:"state"`
	Status string `json:"status"`
	Zone   string `json:"zone"`
	Binary string `json:"binary"`
}

type OsServicePage struct {
	pagination.SinglePageBase
}

func ExtractOsServices(r pagination.Page) ([]OsService, error) {
	var s struct {
		OsServices []OsService `json:"services"`
	}
	err := (r.(OsServicePage)).ExtractInto(&s)
	return s.OsServices, err
}
