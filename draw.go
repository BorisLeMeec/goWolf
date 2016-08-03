package main

import (
	"image/color"
	"math"

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
	var myRay ray
	var fakeScreen []floatPosition
	var prePrePos, prePos, pos floatPosition
	widthImage, _ := img.Size()

	for x := 0; x < widthImage; x++ {
		prePrePos.x = 1
		prePrePos.y = float64(-(widthImage/2)+x) / float64(widthImage)
		prePos.x = prePrePos.x*math.Cos(myData.player.angle) - prePrePos.y*math.Sin(myData.player.angle)
		prePos.y = prePrePos.x*math.Sin(myData.player.angle) + prePrePos.y*math.Cos(myData.player.angle)
		pos.x = prePos.x + myData.player.pos.x*30
		pos.x = prePos.y + myData.player.pos.y*30
		fakeScreen = append(fakeScreen, pos)
		myRay.angle += angleBetweenRay
	}
	for _, value := range fakeScreen {
		setPixel(myData.pix, floatPosToIntPos(value), color.White)
	}
	setPixel(myData.pix, floatPosToIntPos(myData.player.pos), color.White)
}
