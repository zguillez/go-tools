package system

import "fmt"

func Echo(verbose bool, fn func(format string, a ...interface{}), args ...interface{}) {
	if verbose {
		arg := args[1:]
		text := fmt.Sprintf(fmt.Sprintf("%v", args[0]), arg...)
		fn(text)
	}
}
