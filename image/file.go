package image

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"

	"zguillez/go-tools/system"
)

func Load(filepath string) image.Image {
	reader, err := os.Open(filepath)
	system.CheckError(err)
	defer reader.Close()

	bitmap, _, err := image.Decode(reader)
	system.CheckError(err)

	return bitmap
}

func Save(filepath string, bitmap image.Image, level *string) {
	file, err := os.Create(filepath)
	system.CheckError(err)

	quality, err := strconv.Atoi(*level)

	arr := strings.Split(filepath, ".")
	ext := arr[len(arr)-1]

	if ext == "jpg" || ext == "jpeg" {
		color.Green("[encode:jpeg]")
		options := jpeg.Options{Quality: quality}
		jpeg.Encode(file, bitmap, &options)
	} else if ext == "png" {
		color.Green("[encode:png]")
		color.Cyan("[DefaultCompression:0, NoCompression:-1, BestSpeed:-2, BestCompression:-3]")
		if quality > 0 {
			quality = 0
			*level = "0"
		}
		var encoder png.Encoder
		encoder.CompressionLevel = png.CompressionLevel(quality)
		encoder.Encode(file, bitmap)

	} else {
		color.Red("*** error: files extension not suported ***")
		os.Exit(1)
	}
	defer file.Close()
}
