package image

import (
	"image"
	"image/color"
	"image/draw"
)

func ImageToDraw(img image.Image) draw.Image {
	bitmap := image.NewRGBA(img.Bounds())
	draw.Draw(bitmap, bitmap.Bounds(), img, image.Point{}, draw.Src)
	return bitmap
}
func Transparent(width int, height int, fileName string) draw.Image {
	bitmap := image.NewRGBA(image.Rect(0, 0, width, height))
	level := "0"
	Save(fileName, bitmap, &level)
	return bitmap
}
func Solid(width int, height int, fileName string, color color.Color) draw.Image {
	bitmap := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(bitmap, bitmap.Bounds(), &image.Uniform{color}, image.Point{}, draw.Src)
	level := "0"
	Save(fileName, bitmap, &level)
	return bitmap
}
func Overflow(bitmap1 draw.Image, bitmap2 draw.Image, fileName string) draw.Image {
	draw.Draw(bitmap1, bitmap1.Bounds(), bitmap2, image.Point{}, draw.Over)
	level := "0"
	Save(fileName, bitmap1, &level)
	return bitmap1
}
