package system

import (
	"log"

	"github.com/fatih/color"
)

func CheckError(err error) {
	if err != nil {
		color.Red("*** error ***")
		log.Fatal(err)
	}
}
