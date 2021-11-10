package geometry

import (
	"fmt"
	"os"
	"strconv"

	"zguillez.io/gotools/system"
)

func Rectangle() {
	rectangle := Shape{}
	width := system.UserInput{Request: "Insert rectangle width"}
	height := system.UserInput{Request: "Insert rectangle height"}
	complete := make(chan bool)
	go func() {
		width.RequestHandler()
		height.RequestHandler()

		rectangle.Width = ParseFloat(width.Insert)
		rectangle.Height = ParseFloat(height.Insert)
		rectangle.Area = ParseFloat(width.Insert) * ParseFloat(height.Insert)
		rectangle.Perimeter = ParseFloat(width.Insert)*2 + ParseFloat(height.Insert)*2

		complete <- true
	}()

	for <-complete {
		fmt.Println("[output] Rectangle area:", rectangle.Area)
		fmt.Println("[output] Rectangle perimeter:", rectangle.Perimeter)
		os.Exit(1)
	}
}

type Shape struct {
	Width     float64
	Height    float64
	Area      float64
	Perimeter float64
}

func ParseFloat(data string) float64 {
	float, err := strconv.ParseFloat(data, 64)
	if err != nil {
		panic(err)
	}
	return float
}
