package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

func setPixel(img *ebiten.Image, pixels pixelArray, pos position, color color.Color) {
	imageHeight, imageWidth := img.Size()

	if pos.x < 0 || pos.y < 0 || pos.x > uint32(imageWidth-1) || pos.y > uint32(imageHeight-1) {
		return
	}
	index := 4 * (pos.y*uint32(imageWidth) + pos.x)
	r, g, b, a := color.RGBA()
	pixels.pixels[index+0] = uint8(r)
	pixels.pixels[index+1] = uint8(g)
	pixels.pixels[index+2] = uint8(b)
	pixels.pixels[index+3] = uint8(a)
}

func getPixelArray(img *ebiten.Image) (pixelArray, error) {
	var out pixelArray

	imageHeight, imageWidth := img.Size()

	for y := 0; y < imageHeight; y++ {
		for x := 0; x < imageWidth; x++ {
			r, g, b, a := img.At(y, x).RGBA()
			out.pixels = append(out.pixels, uint8(r), uint8(g), uint8(b), uint8(a))
		}
	}
	return out, nil
}

func newPixelArray(width, height uint32) pixelArray {
	var out pixelArray

	out.size.x = width
	out.size.y = height
	for y := uint32(0); y < height; y++ {
		for x := uint32(0); x < width; x++ {
			r, g, b, a := color.Black.RGBA()
			out.pixels = append(out.pixels, uint8(r), uint8(g), uint8(b), uint8(a))
		}
	}
	return out
}

func fill(pixels pixelArray, color color.Color) {
	for x := 0; x < len(pixels.pixels); x += 4 {
		r, g, b, a := color.RGBA()
		pixels.pixels[x] = uint8(r)
		pixels.pixels[x+1] = uint8(g)
		pixels.pixels[x+2] = uint8(b)
		pixels.pixels[x+3] = uint8(a)
	}
}
