package main

type IgnoreList []string

func (ig IgnoreList) Includes(target string) bool {
	for _, name := range ig {
		if target == name {
			return true
		}
	}
	return false
}
