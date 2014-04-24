package main

import "fmt"

type Lang struct {
	Name             string
	Ext              string
	OneLineComment   string
	MultiLineComment string
}

var LangList []*Lang

func init() {
	LangList = []*Lang{
		LangRuby,
	}
}

func DetectLangFromExt(ext string) (*Lang, error) {
	var lang *Lang

	for _, l := range LangList {
		if ext == l.Ext {
			lang = l
			break
		}
	}

	if lang == nil {
		return nil, fmt.Errorf("The file which has \"%s\" extension is not supported.", ext)
	}
	return lang, nil
}
