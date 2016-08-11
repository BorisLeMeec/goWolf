package main

import (
	"image/color"
	"math"
)

func drawWall(pix pixelArray, height, x int) error {
	var pos position
	var myColor color.Color

	if x < 0 || x > int(pix.size.x) {
		return nil
	}
	pos.x = uint32(x)
	if height > int(pix.size.y) {
		height = int(pix.size.y)
	}
	for pos.y = 0; pos.y < pix.size.y; pos.y++ {
		if pos.y < uint32(((int(pix.size.y) - height) / 2)) {
			myColor = color.Black
			setPixel(pix, pos, myColor)
		} else if pos.y < uint32(height+((int(pix.size.y)-height)/2)) {
			myColor = color.White
			setPixel(pix, pos, myColor)
		} else {
			myColor = color.Black
			setPixel(pix, pos, myColor)
		}
	}
	return nil
}

func createFakeScreen() (out []floatPosition) {
	var prePrePos, prePos, newPos floatPosition

	for i := uint32(0); i < myData.pix.size.x; i++ {
		prePrePos.x = 0.5
		prePrePos.y = ((float64(myData.pix.size.x)/2 - float64(i)) / float64(myData.pix.size.x))
		prePos.x = prePrePos.x*math.Cos(myData.player.angle*(math.Pi/180)) - prePrePos.y*math.Sin(myData.player.angle*(math.Pi/180))
		prePos.y = prePrePos.x*math.Sin(myData.player.angle*(math.Pi/180)) + prePrePos.y*math.Cos(myData.player.angle*(math.Pi/180))
		newPos.x = prePos.x + myData.player.pos.x
		newPos.y = prePos.y + myData.player.pos.y
		newPos.x *= float64(15)
		newPos.y *= float64(15)
		setPixel(myData.pix, floatPosToIntPos(newPos), color.White)
		out = append(out, prePrePos)
	}

	return
}

func drawWalls() {
	createFakeScreen()
}
