package system

import (
	"fmt"

	"github.com/fatih/color"
)

func Log(log bool, level int, text string) {
	Logf(log, level, text, []string{}...)
}

func Logf(log bool, level int, text string, args ...string) {

	params := make([]interface{}, len(args))
	for i, param := range args {
		params[i] = param
	}

	if log {
		switch level {
		case 1:
			color.Green(text, params...)
		case 2:
			color.Yellow(text, params...)
		case 3:
			color.Cyan(text, params...)
		case 4:
			color.Red(text, params...)
		default:
			fmt.Println(text)
		}
	}
}
