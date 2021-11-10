package files

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/fatih/color"

	"zguillez.io/gotools/system"
)

func ReadFile(filePath string) {

	if filePath == "" {
		input := system.UserInput{Request: "Insert files path"}
		input.RequestHandler()
		filePath = input.Insert
	}

	output := ""
	complete := make(chan bool)

	go func() {
		output = LoadFile(filePath)
		complete <- true

	}()

	for <-complete {
		color.Cyan("[output] >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		fmt.Println(output)
		color.Cyan("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
		os.Exit(1)
	}
}

func LoadFile(filePath string) string {

	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		return err.Error()
	} else {
		return string(data)
	}
}

func SaveFile(filePath, text string) {

	err := ioutil.WriteFile(filePath, []byte(text), 0644)
	system.CheckError(err)
}

func MoveFile(filePath string, newFilePath string) {

	_, err := os.Stat(filePath)
	system.CheckError(err)

	_, err = os.Stat(newFilePath)
	if err == nil {
		err = os.Remove(newFilePath)
		system.CheckError(err)
	}

	err = os.Rename(filePath, newFilePath)
	system.CheckError(err)
}

func CopyFile(filePath string, newFilePath string) {

	data, err := ioutil.ReadFile(filePath)
	system.CheckError(err)

	err = ioutil.WriteFile(newFilePath, data, 0644)
	system.CheckError(err)
}

func Executable(filePath string) {

	err := os.Chmod(filePath, 0755)
	system.CheckError(err)
}

func DownloadFile(filepath string, url string) {

	resp, err := http.Get(url)
	system.CheckError(err)
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	system.CheckError(err)
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	system.CheckError(err)
}

func CreadeDir(filepath string) {

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		os.Mkdir(filepath, os.ModePerm)
	}

}

func FileExist(filepath string) bool {
	exist := false
	if _, err := os.Stat(filepath); err == nil {
		exist = true
	} else if os.IsNotExist(err) {
		exist = false
	} else {
		system.CheckError(err)
	}
	return exist
}
