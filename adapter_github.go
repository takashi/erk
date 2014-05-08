package main

import (
	"code.google.com/p/goauth2/oauth"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/wsxiaoys/terminal/color"
	"strings"
)

type AdapterGithub struct {
	Token string
}

func (d *AdapterGithub) Update() error {
	var already, newbie int
	var closedIssueTitles []string

	// establish connection with github api by oauth.
	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: config.RemoteConfig.ApiToken},
	}
	client := github.NewClient(t.Client())
	repoInfo := strings.Split(config.RemoteConfig.Repo, "/")

	remoteOpenIssues, remoteClosedIssues, err := GetRemoteIssues(client, repoInfo)
	if err != nil {
		return err
	}

	if len(remoteOpenIssues) != 0 { // if remoteOpenIssues is 0
		for _, i := range IssueList {

			// check if the issue is closed in remote.
			if IsInRemote(remoteClosedIssues, i.Title) {
				closedIssueTitles = append(closedIssueTitles, i.Title)
				already++
			} else {
				frag := fmt.Sprintf("filename: %s\nline: %d\nlabel: %s\n\n```%s\n%s\n```\n", i.FilePath, i.Line, i.Label, i.Lang, i.Fragment)
				labels := []string{"erk", i.Label}
				_, _, err := client.Issues.Create(repoInfo[0], repoInfo[1], &github.IssueRequest{Title: &i.Title, Body: &frag, Labels: labels})
				if err != nil {
					return err
				}
				newbie++
			}
		}
	} else { // if there is remoteOpenIssues
		for _, i := range IssueList {
			// check if the issue is closed in remote.
			if IsInRemote(remoteClosedIssues, i.Title) {
				closedIssueTitles = append(closedIssueTitles, i.Title)
			} else {
				roi := GetInRemoteIssueByTitle(remoteOpenIssues, i.Title)
				if roi != nil {
					already++
				} else {
					frag := fmt.Sprintf("filename: %s\nline: %d\nlabel: %s\n\n```%s\n%s\n```\n", i.FilePath, i.Line, i.Label, i.Lang, i.Fragment)
					labels := []string{"erk", i.Label}
					_, _, err := client.Issues.Create(repoInfo[0], repoInfo[1], &github.IssueRequest{Title: &i.Title, Body: &frag, Labels: labels})
					if err != nil {
						return err
					}
					newbie++
				}
			}
		}

	}

	fmt.Printf("%d new issues are registerd in remote, %d issues is already registerd. \n", newbie, already)

	// print closed issues.
	if len(closedIssueTitles) != 0 {
		fmt.Printf("\nand the issues: \n")
		for _, t := range closedIssueTitles {
			color.Printf("- @{!w}%s\n", t)
		}
		color.Printf("seems to be closed. Do you forget to delete these \"%s\" line?\n", config.Label)
	}
	return nil
}

func IsInRemote(issues []github.Issue, title string) bool {
	for _, i := range issues {
		if *i.Title == title {
			return true
		}
	}
	return false
}

func GetInRemoteIssueByTitle(issues []github.Issue, title string) *github.Issue {
	for _, i := range issues {
		if *i.Title == title {
			return &i
		}
	}
	return nil
}

func GetRemoteIssues(client *github.Client, repoInfo []string) (remoteOpenIssues []github.Issue, remoteClosedIssues []github.Issue, err error) {
	// get all remote issues that are opened.
	remoteOpenIssues, _, err = client.Issues.ListByRepo(repoInfo[0], repoInfo[1], &github.IssueListByRepoOptions{Labels: []string{"erk"}})
	// get all remote issues that are closed.
	remoteClosedIssues, _, err = client.Issues.ListByRepo(repoInfo[0], repoInfo[1], &github.IssueListByRepoOptions{Labels: []string{"erk"}, State: "closed"})
	return
}
