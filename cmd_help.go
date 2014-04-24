package main

import (
	"bytes"
	"fmt"
)

// add space to line up
func Ljust(word string, length int) string {
	if len(word) > length {
		return word
	} else {
		var buffer bytes.Buffer
		l := length - len(word)
		buffer.WriteString(word)
		for i := 0; l > i; i++ {
			buffer.WriteString(" ")
		}
		return buffer.String()
	}
}

var CmdHelp = &Cmd{
	Name: "help",
	Desc: "show help.",
	Fn: func(args []string) error {
		fmt.Printf(`erk v%s
Usage: erk commands are:

`, VERSION)
		maxLen := func() int {
			max := 0
			for _, cmd := range CmdList {
				if max < len(cmd.Name) {
					max = len(cmd.Name)
				}
			}
			return max
		}()
		for _, cmd := range CmdList {
			fmt.Printf("  "+Ljust(cmd.Name, maxLen)+"   %s\n", cmd.Desc)
		}
		return nil
	},
}
