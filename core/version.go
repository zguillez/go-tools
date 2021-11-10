package core

import (
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"

	"zguillez/go-tools/files"
	"zguillez/go-tools/text"
)

func Version(level string) {

	versions := map[string]bool{
		"minor": true,
		"major": true,
		"patch": true,
	}
	if versions[level] {

		paquete := readPackageFile()
		version := strings.Split(paquete.Version, ".")
		if len(version) != 3 {
			color.Red("*** error *** Invalid packaje version %v", paquete.Version)
			os.Exit(1)
		}
		paqueteText := files.LoadFile("./package.json")

		color.Cyan("package: %v", paquete.Version)

		var i int
		if level == "major" {
			i = 0
			version[1] = "0"
			version[2] = "0"
		} else if level == "minor" {
			i = 1
			version[2] = "0"
		} else if level == "patch" {
			i = 2
		}
		n, _ := strconv.Atoi(version[i])
		version[i] = strconv.Itoa(n + 1)
		versionActual := strings.Join(version, ".")

		color.Cyan("%v update to: %v", level, versionActual)

		paqueteTexto := text.Replace(paqueteText, paquete.Version, versionActual, -1)
		files.SaveFile("./package.json", paqueteTexto)

	} else {
		color.Red("*** error *** Invalid version -l (minor|major|patch)")
		os.Exit(1)
	}
}
