package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

func ParseFiles(dir string) error {
	err := filepath.Walk(dir, WalkFn)
	if err != nil {
		return err
	}
	return nil
}

func WalkFn(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	isIgnored := config.Ignore.Includes(info.Name())
	if isIgnored && info.IsDir() {
		// ignored directory
		return filepath.SkipDir // http://play.golang.org/p/Qc4-NIs25c
	} else if isIgnored {
		// ignored file.
	} else if info.IsDir() {
		// do nothing
	} else {
		ParseIssues(path string, info os.FileInfo)
	}
	return nil
}

func ParseIssues() {
	// do the file parse here
	lang := DetectLangFromExt(filepath.Ext(path))
	if lang != nil {
		body, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		re, _ := regexp.Compile(`# \[TODO\].*`)
		one := re.Find(body)
		if one != nil {
			fmt.Printf("Find a TODO: %s \n    at %s\n", one, path)
		}
	}
}
