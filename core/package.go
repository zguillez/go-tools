package core

import (
	"encoding/json"
	"io/ioutil"

	"github.com/fatih/color"

	"zguillez.io/gotools/system"
)

type PackageJson struct {
	Version string
}

func Package() {
	color.Green(readPackageFile().Version)
}

func readPackageFile() PackageJson {
	raw, err := ioutil.ReadFile("./package.json")
	system.CheckError(err)

	var paquete PackageJson
	json.Unmarshal(raw, &paquete)

	return paquete
}
