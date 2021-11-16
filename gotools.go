package gotools

import (
	"github.com/fatih/color"
	"github.com/zguillez/go-tools/core"
	"github.com/zguillez/go-tools/files"
	"github.com/zguillez/go-tools/geometry"
	"github.com/zguillez/go-tools/image"
	"github.com/zguillez/go-tools/system"
	"github.com/zguillez/go-tools/text"
)

func Help() {
	core.Help()
}
func Version() {
	color.Green("v0.1.36")
}

func Command(command string, verbose bool, args ...string) {
	system.Echo(verbose, color.Yellow, "[command: %v]", command)
	switch command {
	case "rectangle":
		geometry.Rectangle()
	case "replace":
		text.ReplaceCli()
	case "readfile":
		files.ReadFile(args[0])
	case "minimize":
		image.Minimize(args[0], args[1], args[2])
	case "version":
		core.Version(args[0])
	default:
		system.Echo(verbose, color.Red, "[error: command %v unknown]", command)
	}
}
