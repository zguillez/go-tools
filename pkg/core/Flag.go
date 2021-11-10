package core

import (
	"flag"
)

func FlagString(tags []string, value string, usage string) *string {
	tag := &value
	for i, name := range tags {
		if i == 0 {
			tag = flag.String(name, value, usage)
		} else {
			flag.StringVar(tag, name, value, usage)
		}
	}
	return tag
}
func FlagBool(tags []string, value bool, usage string) *bool {
	tag := &value
	for i, name := range tags {
		if i == 0 {
			tag = flag.Bool(name, value, usage)
		} else {
			flag.BoolVar(tag, name, value, usage)
		}
	}
	return tag
}
func FlagParse() {
	flag.Parse()
}
