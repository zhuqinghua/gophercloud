package agents

import "github.com/zhuqinghua/gophercloud.git/pagination"

type Agent struct {
	AdminStateUp     bool   `json:"admin_state_up"`
	Alive            bool   `json:"alive"`
	AvailabilityZone string `json:"availability_zone"`
	Binary           string `json:"binary"`
	Host             string `json:"host"`
	ID               string `json:"id"`
	Type             string `json:"agent_type"`
}

type AgentPage struct {
	pagination.SinglePageBase
}

func ExtractAgents(r pagination.Page) ([]Agent, error) {
	var s struct {
		Agents []Agent `json:"agents"`
	}
	err := (r.(AgentPage)).ExtractInto(&s)
	return s.Agents, err
}
