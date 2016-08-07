package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

func drawWall(img *ebiten.Image, pix pixelArray, height, x int) error {
	var pos position
	var myColor color.Color
	imageHeight, imageWidth := img.Size()

	if x < 0 || x > imageWidth {
		return nil
	}
	pos.x = uint32(x)
	if height > imageHeight {
		height = imageHeight
	}
	for pos.y = 0; pos.y < uint32(imageHeight); pos.y++ {
		if pos.y < uint32(((imageHeight - height) / 2)) {
			myColor = color.Black
			setPixel(pix, pos, myColor)
		} else if pos.y < uint32(height+((imageHeight-height)/2)) {
			myColor = color.White
			setPixel(pix, pos, myColor)
		} else {
			myColor = color.Black
			setPixel(pix, pos, myColor)
		}
	}
	return nil
}

func drawScreen(img *ebiten.Image, myData data) {
	drawMiniMap(myData)
	blit(myData.pix, myData.miniMap.pix, myData.miniMap.posStart, myData.miniMap.pix.size)
}
