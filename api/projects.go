package api

import (
	"log"

	"github.com/jaymickey/circleops/client"
)

type Project struct {
	Name         string `json:"reponame"`
	Organisation string `json:"username"`
	Language     string `json:"language"`
}

func GetProjects(client *client.Client) []*Project {
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	projects := []*Project{}
	if err := client.SetHeaders(headers).SetMethod("Get").Run(&projects); err != nil {
		log.Fatalf("error getting projects: %v", err)
	}

	return projects
}
