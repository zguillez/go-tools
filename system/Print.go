package system

import "github.com/fatih/color"

func Print(verbose bool, col color.Attribute, text string) {
	if verbose {
		c := color.New(col)
		c.Println(text)
	}
}
