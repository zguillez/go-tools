package core

import "github.com/fatih/color"

func Help() {
	color.Yellow("[gotools help]")

	color.Cyan("[update package version]")
	color.Cyan("\timport \"github.com/zguillez/go-tools/core\"\n")
	color.Cyan("\tcore.Version(\"patch\") //or minor | major\n")
	color.Cyan("\n")

	color.Cyan("[minimize files]")
	color.Cyan("\timport \"github.com/zguillez/go-tools/image\"\n")
	color.Cyan("\tinput := \"img.png\"\n")
	color.Cyan("\toutput := \"img.min.png\"\n")
	color.Cyan("\tlevel := \"-2\"\n")
	color.Cyan("\t\t0 - DefaultCompression")
	color.Cyan("\t\t-1 - NoCompression")
	color.Cyan("\t\t-2 - BestSpeed")
	color.Cyan("\t\t-3 - BestCompression")
	color.Cyan("\t\t(1-100) of JPG")
	color.Cyan("\timage.Minimize(input, output, level)\n")
	color.Cyan("\n")
}
