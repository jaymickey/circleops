package api

import (
	"github.com/jaymickey/circleops/client"
)

type Organisation struct {
	Name    string `json:"name"`
	IsAdmin bool   `json:"admin"`
	VCS     string `json:"vcs_type"`
}

// https://circleci.com/api/v1.1/user/organizations
// {
// 	"org": true,
// 	"avatar_url": "https://avatars2.githubusercontent.com/u/5773286?v=4",
// 	"admin": true,
// 	"name": "conde-nast-international",
// 	"piggieback_org_maps": [],
// 	"piggieback_orgs": [],
// 	"vcs_type": "github",
// 	"login": "conde-nast-international",
// 	"num_paid_containers": 1
// }

//
func GetOrganisations(client *client.Client, includeUser bool) ([]*Organisation, error) {
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	if includeUser {
		client.AddQuery("include-user", includeUser)
	}
	var orgs []*Organisation

	if err := client.SetHeaders(headers).SetMethod("GET").Run(&orgs); err != nil {
		return nil, err
	}

	return orgs, nil
}
