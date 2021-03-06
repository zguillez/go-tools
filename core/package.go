package core

import (
	"encoding/json"
	"io/ioutil"

	"github.com/fatih/color"

	"github.com/zguillez/go-tools/system"
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
