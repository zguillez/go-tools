package text

import (
	"fmt"
	"os"
	"strings"

	"zguillez/go-tools/system"
)

func Replace(input string, search string, replace string, count int) string {

	return strings.Replace(input, search, replace, count)
}

func ReplaceCli() {

	input := system.UserInput{Request: "Insert text"}
	search := system.UserInput{Request: "Search"}
	replace := system.UserInput{Request: "Replace"}
	output := ""
	complete := make(chan bool)
	go func() {
		input.RequestHandler()
		search.RequestHandler()
		replace.RequestHandler()

		output = Replace(input.Insert, search.Insert, replace.Insert, -1)
		complete <- true
	}()

	for <-complete {
		fmt.Println("[output] ", output)
		os.Exit(1)
	}
}
