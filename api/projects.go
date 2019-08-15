package api

import (
	"mickey.dev/circleops/client"
)

type Project struct {
	Name         string `json:"reponame"`
	Organisation string `json:"username"`
	Language     string `json:"language"`
}

func GetProjects(client *client.Client) ([]*Project, error) {
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	projects := []*Project{}
	if err := client.SetHeaders(headers).SetMethod("Get").Run(&projects); err != nil {
		return nil, err
	}

	return projects, nil
}
