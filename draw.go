package main

import (
	"image/color"
	"math"

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
			myColor = color.Black
			setPixel(img, pixels, pos, myColor)
		} else if pos.y < uint32(height+((imageHeight-height)/2)) {
			myColor = color.White
			setPixel(img, pixels, pos, myColor)
		} else {
			myColor = color.Black
			setPixel(img, pixels, pos, myColor)
		}
	}
	return nil
}

func drawScreen(img *ebiten.Image, myData data) {
	var myRay ray
	var fakeScreen []floatPosition
	var pos, prePos, prePrePos floatPosition
	widthImage, _ := img.Size()

	for x := 0; x < widthImage; x++ {
		prePrePos.x = 1
		prePrePos.y = float64(((widthImage / 2) + x) / widthImage)
		prePos.x = prePrePos.x*math.Cos(myData.player.angle) - prePrePos.y*math.Sin(myData.player.angle)
		prePos.y = prePrePos.x*math.Sin(myData.player.angle) + prePrePos.y*math.Cos(myData.player.angle)
		pos.x = prePos.x + myData.player.pos.x
		pos.x = prePos.y + myData.player.pos.y
		fakeScreen = append(fakeScreen, pos)
		myRay.angle += angleBetweenRay
	}
	// fmt.Println(fakeScreen)
	// os.Exit(4)
	setPixel(img, myData.pixelArray, floatPosToIntPos(myData.player.pos), color.White)
}
