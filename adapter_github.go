package main

import (
	"code.google.com/p/goauth2/oauth"
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
	issues, _, err := client.Issues.ListByRepo(repoInfo[0], repoInfo[1], nil)
	if err != nil {
		return err
	}
	for _, issue := range issues {
		println(*issue.Title)
	}
	return nil
}
