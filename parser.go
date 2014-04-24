package main

import (
	"os"
	"path/filepath"
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
		// do the file parse here
		println(info.Name())
	}
	return nil
}
