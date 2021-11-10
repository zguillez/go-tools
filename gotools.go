package gotools

import (
	"github.com/fatih/color"
	"zguillez/go-tools/core"
	"zguillez/go-tools/files"
	"zguillez/go-tools/geometry"
	"zguillez/go-tools/image"
	"zguillez/go-tools/text"
)

func main() {
	version := core.FlagBool([]string{"version", "v"}, false, "Package version")
	help := core.FlagBool([]string{"help", "h"}, false, "Help")
	command := core.FlagString([]string{"command", "c"}, "", "Command to execute:\ntools -c version -l minor")
	input := core.FlagString([]string{"input", "i"}, "", "Input value text")
	output := core.FlagString([]string{"output", "o"}, "", "Output value text")
	level := core.FlagString([]string{"level", "l"}, "100", "Level value int")
	core.FlagParse()

	if *help {
		core.Help()
	} else if *version {
		// core.Package()
		color.Green("v0.2.0")
	} else if *command != "" {
		color.Yellow("[command: %v]", *command)
		switch *command {
		case "rectangle":
			geometry.Rectangle()
		case "replace":
			text.ReplaceCli()
		case "readfile":
			files.ReadFile(*input)
		case "minimize":
			image.Minimize(*input, *output, *level)
		case "version":
			core.Version(*level)
		default:
			color.Red("[error: command %v unknown]", *command)
		}
	} else {
		color.Red("[error: tool -c command empty]")
	}

}

func Hello() string {
	return "HELLO!!!!!"
}
