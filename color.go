package main

import "image/color"

var (
	red    = color.RGBA{255, 0, 0, 255}
	green  = color.RGBA{0, 255, 0, 255}
	blue   = color.RGBA{0, 0, 255, 255}
	yellow = color.RGBA{255, 215, 0, 255}
	pink   = color.RGBA{255, 20, 147, 255}
)

func (pix *pixelArray) getColorAt(index uint32) (r, g, b, a uint8) {
	r = pix.pixels[index+0]
	g = pix.pixels[index+1]
	b = pix.pixels[index+2]
	a = pix.pixels[index+3]
	return
}

func (pix *pixelArray) setColorAt(index uint32, r, g, b, a uint8) {
	pix.pixels[index+0] = r
	pix.pixels[index+1] = g
	pix.pixels[index+2] = b
	pix.pixels[index+3] = a
}

func (pix *pixelArray) changeOpacity(opacity float32) {
	a := uint8(opacity * 255)

	for x := uint32(0); x < pix.size.x*pix.size.y*4; x += 4 {
		pix.pixels[x+3] = a
	}
}
