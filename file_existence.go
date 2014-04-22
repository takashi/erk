package main

import "os"

// http://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go
func CheckFileExistence(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}
