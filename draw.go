package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

func drawWall(img *ebiten.Image, pixels []uint8, height, x int) error {
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
			fmt.Printf("black : %d\n", pos.y)
			myColor = color.Black
			setPixel(img, pixels, pos, myColor)
		} else if pos.y < uint32(height+((imageHeight-height)/2)) {
			fmt.Printf("white : %d\n", pos.y)
			myColor = color.White
			setPixel(img, pixels, pos, myColor)
		} else {
			fmt.Printf("black : %d\n", pos.y)
			myColor = color.Black
			setPixel(img, pixels, pos, myColor)
		}
	}
	return nil
}
