package main

import (
	"code.google.com/p/goauth2/oauth"
	"fmt"
	"github.com/google/go-github/github"
	"strings"
)

type AdapterGithub struct {
	Token string
}

func (d *AdapterGithub) Update() error {
	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: config.RemoteConfig.ApiToken},
	}
	client := github.NewClient(t.Client())
	repoInfo := strings.Split(config.RemoteConfig.Repo, "/")
	remoteIssues, _, err := client.Issues.ListByRepo(repoInfo[0], repoInfo[1], &github.IssueListByRepoOptions{Labels: []string{"erk"}})
	for _, issue := range IssueList {
		frag := fmt.Sprintf("filename: %s\nline: %d\nlabel: %s\n\n```go\n%s\n```\n", issue.FilePath, issue.Line, issue.Label, issue.Fragment)
		labels := []string{"erk", issue.Label}
		client.Issues.Create(repoInfo[0], repoInfo[1], &github.IssueRequest{Title: &issue.Title, Body: &frag, Labels: labels})
	}
	return nil
}
