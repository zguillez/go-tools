package system

import "fmt"

func Echo(verbose bool, fn func(format string, a ...interface{}), arg ...interface{}, ) {
	if verbose {
		arg_ := arg[1:]
		text := fmt.Sprintf(fmt.Sprintf("%v", arg[0]), arg_...)
		fn(text)
	}
}
