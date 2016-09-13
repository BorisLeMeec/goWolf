package main

import (
	"image/color"
	"math"
	"sort"
)

func drawWall(pix PixelArray, height, x int) error {
	var pos position
	var myColor color.Color

	if x < 0 || x > int(pix.size.x) {
		return nil
	}
	pos.x = x
	if height > int(pix.size.y) {
		height = int(pix.size.y)
	}
	for pos.y = 0; pos.y < pix.size.y; pos.y++ {
		if pos.y < ((int(pix.size.y) - height) / 2) {
			myColor = color.Black
			pix.SetPixel(pos, myColor)
		} else if pos.y < height+((int(pix.size.y)-height)/2) {
			myColor = color.White
			pix.SetPixel(pos, myColor)
		} else {
			myColor = color.Black
			pix.SetPixel(pos, myColor)
		}
	}
	return nil
}

func createFakeScreen() (out []floatPosition) {
	var prePrePos, prePos, newPos floatPosition

	for i := 0; i < myData.pix.size.x; i++ {
		prePrePos.x = 0.5
		prePrePos.y = ((float64(myData.pix.size.x)/2 - float64(i)) / float64(myData.pix.size.x))
		prePos.x = prePrePos.x*math.Cos(myData.player.angle*(math.Pi/180)) - prePrePos.y*math.Sin(myData.player.angle*(math.Pi/180))
		prePos.y = prePrePos.x*math.Sin(myData.player.angle*(math.Pi/180)) + prePrePos.y*math.Cos(myData.player.angle*(math.Pi/180))
		newPos.x = prePos.x + myData.player.pos.x
		newPos.y = prePos.y + myData.player.pos.y
		newPos.x *= float64(15)
		newPos.y *= float64(15)
		myData.pix.SetPixel(newPos.toIntPos(), color.White)
		out = append(out, prePrePos)
	}
	return
}

func drawWalls() {
	fakeScreen := createFakeScreen()
	var vectors = make([]vect, len(fakeScreen))
	var k []float64
	// var index uint32

	for x := 0; x < len(fakeScreen); x++ {
		k = make([]float64, myData.theMap.size.x+myData.theMap.size.y)
		vectors[x] = vect{myData.player.pos.x - fakeScreen[x].x, myData.player.pos.y - fakeScreen[x].x}
		for i := 0; i < myData.theMap.size.x; i++ {
			k[i] = (float64(i) - myData.player.pos.x) / vectors[x].x
		}
		for i := myData.theMap.size.x; i < myData.theMap.size.x+myData.theMap.size.y; i++ {
			k[i] = (float64(i) - myData.player.pos.y) / vectors[x].y
		}
		sort.Sort(sort.Float64Slice(k))
		i := 0
		for ; k[i] <= 0; i++ {
		}
	}
}
