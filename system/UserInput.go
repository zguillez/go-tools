package system

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type UserInput struct {
	Request string
	Insert  string
}

func (req *UserInput) RequestHandler() {
	fmt.Printf("[input] " + req.Request + ": ")
	reader := bufio.NewReader(os.Stdin)
	insert, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	} else {
		req.Insert = strings.Trim(insert, " \n")
	}
}
