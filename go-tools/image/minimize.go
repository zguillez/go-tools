package image

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func Minimize(input string, output string, level string) {

	if input == "" {
		color.Red("*** error: -i input files path is empty ***")
		os.Exit(1)
	}
	if output == "" {
		arr := strings.Split(input, ".")
		output = arr[0] + ".min." + arr[1]
	}

	bitmap := Load(input)
	Save(output, bitmap, &level)

	fmt.Printf("[quality:%s] %s\n", level, output)
}
