package main

type Issue struct {
	Title    string
	FileName string
	FilePath string
	Fragment string
	Label    string
	Line     int
	Md5      string
}

var IssueList []*Issue
