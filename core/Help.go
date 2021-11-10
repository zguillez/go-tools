package core

import "github.com/fatih/color"

func Help() {
	color.Yellow("[go-tools help]")

	color.Cyan("[update package version]")
	color.Cyan("\tgo-tools -c version -l patch\n")
	color.Cyan("\n")

	color.Cyan("[minimize files]")
	color.Cyan("\tgo-tools -c minimize -i logo1.png -l 0")
	color.Cyan("\t\t-l 0 - DefaultCompression")
	color.Cyan("\t\t-l -1 - NoCompression")
	color.Cyan("\t\t-l -2 - BestSpeed")
	color.Cyan("\t\t-l -3 - BestCompression")
	color.Cyan("\tgo-tools -c minimize -i logo1.jpg -l 80")
	color.Cyan("\t\t-l [0-100]")
	color.Cyan("\n")
}
