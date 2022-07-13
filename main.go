package main

import (
	"context"

	"github.com/flosch/pongo2/v6"
	"github.com/google/go-github/v45/github"
)

/* -------------------------------------------------------------------------- */
/*                           通过 git issue 生成 readme                       */
/* -------------------------------------------------------------------------- */

func main() {
	issues, err := getAllIssues()
	if err != nil {
		panic(err)
	}
	output := parseIssueTitleAndLink(issues)
	t, _ := pongo2.FromString("")
	t.Execute(pongo2.Context{
		"arr": output,
	})
}

// GetAllIssues 拉取所有 issues
func getAllIssues() ([]*github.Issue, error) {
	client := github.NewClient(nil)
	ctx := context.Background()
	issues, _, err := client.Issues.List(ctx, true, nil)
	return issues, err
}

// 用于渲染
type issue struct {
	Title string `json:"title"`
	Link  string `json:"Link"`
}

// 解析所有 issue 标题和链接
func parseIssueTitleAndLink(issues []*github.Issue) []*issue {
	res := make([]*issue, len(issues))
	for i, v := range issues {
		if v.URL == nil || v.Title == nil {
			continue
		}
		res[i] = &issue{
			Title: *v.Title,
			Link:  *v.URL,
		}
	}
	return res
}

func renderReadme() error {
}
