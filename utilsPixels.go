package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

func setPixel(img *ebiten.Image, pixels []uint8, pos position, color color.Color) {
	imageHeight, imageWidth := img.Size()

	if pos.x < 0 || pos.y < 0 || pos.x > uint32(imageWidth-1) || pos.y > uint32(imageHeight-1) {
		return
	}
	index := 4 * (pos.y*uint32(imageWidth) + pos.x)
	r, g, b, a := color.RGBA()
	pixels[index+0] = uint8(r)
	pixels[index+1] = uint8(g)
	pixels[index+2] = uint8(b)
	pixels[index+3] = uint8(a)
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

func newPixelArray(width, height int) []uint8 {
	var out []uint8

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, a := color.Black.RGBA()
			out = append(out, uint8(r), uint8(g), uint8(b), uint8(a))
		}
	}
	return out
}

func fill(pixelArray []uint8, color color.Color) {
	for x := 0; x < len(pixelArray); x += 4 {
		r, g, b, a := color.RGBA()
		pixelArray[x] = uint8(r)
		pixelArray[x+1] = uint8(g)
		pixelArray[x+2] = uint8(b)
		pixelArray[x+3] = uint8(a)
	}
}
