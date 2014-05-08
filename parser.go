package main

import (
	"bufio"
	"github.com/takashi/erk/language"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func ParseFiles( /*dir string*/) error {
	err := filepath.Walk("./", WalkFn)
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
		err := ParseIssues(path, info)
		if err != nil {
			return err
		}
	}
	return nil
}

func ParseIssues(path string, info os.FileInfo) error {
	// do the file parse here
	lang := language.DetectLangFromExt(filepath.Ext(path))
	if lang != nil {
		err := ParseIssuesByLang(lang, path, info)
		if err != nil {
			return err
		}
	}
	return nil
}

func ParseIssuesByLang(lang *language.Lang, path string, info os.FileInfo) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	scanner := bufio.NewScanner(file)
	line := 1
	re, _ := regexp.Compile(lang.OneLineComment + `.*\[(?i)` + config.Label + `\].*`)
	commnentLabelRe, _ := regexp.Compile(lang.OneLineComment + `.*\[(?i)` + config.Label + `\].*\-`)
	for scanner.Scan() {
		one := re.Find([]byte(scanner.Text()))
		if one != nil {
			title := strings.TrimPrefix(string(commnentLabelRe.ReplaceAll(one, []byte(""))), " ")
			frag := scanner.Text() + "\n"

			// copy scanner instance
			s := *scanner
			i := 0
			// get 10 line under the todo.
			for s.Scan() {
				if i > 10 {
					break
				}
				frag += (s.Text() + "\n")
				i++
			}

			issue := &Issue{
				Title:    title,
				FileName: info.Name(),
				Label:    config.Label,
				Line:     line,
				FilePath: path,
				Fragment: frag,
				Lang:     lang.Name}
			IssueList = append(IssueList, issue)
		}
		line += 1
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
