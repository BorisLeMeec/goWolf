package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

func setPixel(img *ebiten.Image, pixels []uint8, pos position, color color.Color) {
	imageHeight, imageWidth := img.Size()
	if pos.x < 0 || pos.y < 0 || pos.x > uint32(imageWidth) || pos.y > uint32(imageHeight) {
		return
	}
	index := 4 * (pos.y*uint32(imageWidth) + pos.x)
	pixels[index+0] = 255
	pixels[index+1] = 255
	pixels[index+2] = 255
	pixels[index+3] = 255
}

func getPixelArray(img *ebiten.Image) ([]uint8, error) {
	var out []uint8

	imageHeight, imageWidth := img.Size()
	for y := 0; y < imageHeight; y++ {
		for x := 0; x < imageWidth; x++ {
			r, g, b, a := img.At(y, x).RGBA()
			out = append(out, uint8(r), uint8(g), uint8(b), uint8(a))
		}
	}
	return out, nil
}
