package server

import (
	"errors"

	"github.com/kloudlite/kl/domain/client"
	fn "github.com/kloudlite/kl/pkg/functions"
	"github.com/kloudlite/kl/pkg/ui/fzf"
)

type Project struct {
	DisplayName string   `json:"displayName"`
	Metadata    Metadata `json:"metadata"`
	Status      Status   `json:"status"`
}

type ProjectList struct {
	Edges Edges[Project] `json:"edges"`
}

func ListProjects(options ...fn.Option) ([]Project, error) {
	accountName := fn.GetOption(options, "accountName")
	clusterName := fn.GetOption(options, "clusterName")

	var err error

	if accountName == "" {
		accountName, err = client.CurrentAccountName()

		if err != nil {
			return nil, err
		}
	}

	if clusterName == "" {
		clusterName, err = client.CurrentClusterName()

		if err != nil {
			return nil, err
		}
	}

	cookie, err := getCookie()
	if err != nil {
		return nil, err
	}

	respData, err := klFetch("cli_listProjects", map[string]any{
		"pq": map[string]any{
			"orderBy":       "name",
			"sortDirection": "ASC",
			"first":         99999999,
		},
	}, &cookie)

	if err != nil {
		return nil, err
	}

	if fromResp, err := GetFromRespForEdge[Project](respData); err != nil {
		return nil, err
	} else {
		return fromResp, nil
	}
}

func SelectProject(projectName string) (*Project, error) {
	projects, err := ListProjects()
	if err != nil {
		return nil, err
	}

	if projectName != "" {
		for _, a := range projects {
			if a.Metadata.Name == projectName {
				return &a, nil
			}
		}
		return nil, errors.New("you don't have access to this account")
	}

	project, err := fzf.FindOne(
		projects,
		func(project Project) string {
			return project.DisplayName
		},
		fzf.WithPrompt("Select Project > "),
	)

	if err != nil {
		return nil, err
	}

	return project, nil
}
