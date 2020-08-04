package main

import (
	"image/color"
)

var teal color.Color = color.RGBA{0, 200, 200, 255}
var red color.Color = color.RGBA{200, 30, 30, 255}
var blue color.Color = color.RGBA{0, 0, 255, 255}

func main() {
	DrawRectX("myimage.png", 500, teal, red, blue)

	DrawShessTable("shesstable.png", 240)

}
