package gotools

import (
	"github.com/fatih/color"
	"github.com/zguillez/go-tools/core"
)

func Help() {
	core.Help()
}
func Version() {
	color.Green("v0.1.24")
}
