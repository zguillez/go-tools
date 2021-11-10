package xml

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"zguillez/go-tools/system"
)

func Get(url string) []byte {
	resp, err := http.Get(url)
	system.CheckError(err)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		system.CheckError(fmt.Errorf("Status error: %v", resp.StatusCode))
	}
	data, err := ioutil.ReadAll(resp.Body)
	system.CheckError(err)

	return data
}

func ForceGet(url string) []byte {
	resp, err := http.Get(url)
	system.CheckError(err)
	defer resp.Body.Close()

	var data []byte
	if resp.StatusCode == http.StatusOK {
		data, err = ioutil.ReadAll(resp.Body)
	}
	system.CheckError(err)

	return data
}
