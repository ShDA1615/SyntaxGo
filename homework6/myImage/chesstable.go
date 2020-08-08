package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func DrawShessTable(name string, size int) {
	file, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img := image.NewRGBA(image.Rect(0, 0, size, size))
	colors := make(map[int]color.RGBA, 2)

	colors[0] = color.RGBA{0, 0, 0, 255}
	colors[1] = color.RGBA{255, 255, 255, 255}

	color := 0

	startX := 0

	for x := 0; x < 8; x++ {

		startY := 0
		for y := 0; y < 8; y++ {

			draw.Draw(img, image.Rect(startX, startY, startX+int(size/8), startY+int(size/8)),
				&image.Uniform{colors[color]}, image.ZP, draw.Src)

			startY += int(size / 8)
			color = 1 - color
		}
		startX += int(size / 8)
		color = 1 - color
	}

	png.Encode(file, img)
}
