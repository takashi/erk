package language

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
		LangJs,
		LangGo,
	}
}

func DetectLangFromExt(ext string) *Lang {
	var lang *Lang

	for _, l := range LangList {
		if ext == l.Ext {
			lang = l
			break
		}
	}

	if lang == nil {
		return nil
	}
	return lang
}
