package main

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/flosch/pongo2/v6"
	"github.com/google/go-github/v45/github"
	"golang.org/x/oauth2"
)

/* ------------------------- 通过 git issue 生成 readme ------------------------- */

const readmeTemplate = `# Git Blog

## Top
{%- for post in allList %}
{%- if post.IsTop %} 
- [{{ post.Title }}]({{ post.Link }}) 
{%- endif %} 
{%- endfor %}

## All
{%- for post in allList %} 
- [{{ post.Title }}]({{ post.Link }}) 
{%- endfor %}
`

var GITHUB_TOKEN = os.Getenv("GITHUB_TOKEN")

func main() {
	issues, err := getAllIssues()
	if err != nil {
		panic(err)
	}
	output := parseIssueTitleAndLink(issues)
	t := pongo2.Must(pongo2.FromString(readmeTemplate))
	newReadMe, err := t.Execute(pongo2.Context{
		"allList": output,
	})
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("README.md", []byte(newReadMe), 0666)
}

func getAllIssues() ([]*github.Issue, error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: GITHUB_TOKEN},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	issues, _, err := client.Issues.ListByRepo(ctx, "linbuxiao", "gitblog", nil)
	return issues, err
}

type issue struct {
	Title string `json:"title"`
	Link  string `json:"link"`
	IsTop bool   `json:"is_top"`
}

func parseIssueTitleAndLink(issues []*github.Issue) []*issue {
	res := make([]*issue, len(issues))
	for i, v := range issues {
		if v.URL == nil || v.Title == nil {
			continue
		}
		isTop := false
		for _, x := range v.Labels {
			if x.GetName() == "Top" {
				isTop = true
				break
			}
		}
		res[i] = &issue{
			Title: *v.Title,
			Link:  *v.URL,
			IsTop: isTop,
		}
	}
	return res
}
