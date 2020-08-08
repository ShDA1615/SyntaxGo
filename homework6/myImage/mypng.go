package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func DrawRectX(name string, size int, back, line1, line2 color.Color) {
	file, err := os.Create(name)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	img := image.NewRGBA(image.Rect(0, 0, size, size))

	draw.Draw(img, img.Bounds(), &image.Uniform{teal}, image.ZP, draw.Src)
	for x := 0; x < size; x++ {
		y := x
		img.Set(x, y, red)
		x1 := x
		y1 := size - y
		img.Set(x1, y1, blue)
	}

	png.Encode(file, img)
}
