package main

import (
	"fmt"
	"image/color"
	"math"
	"os"

	"github.com/hajimehoshi/ebiten"
)

var myData data
var height = 400
var width = 400
var angleBetweenRay = 60 / float64(width)

func checkKey() {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		myData.player.pos.y += math.Sin(myData.player.angle * (math.Pi / 180))
		myData.player.pos.x += math.Cos(myData.player.angle * (math.Pi / 180))
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		myData.player.pos.y -= math.Sin(myData.player.angle * (math.Pi / 180))
		myData.player.pos.x -= math.Cos(myData.player.angle * (math.Pi / 180))
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		myData.player.angle -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		myData.player.angle += 5
	}
	if myData.player.angle > 360 {
		myData.player.angle = 0
	}
	if myData.player.angle < 0 {
		myData.player.angle = 360
	}
}

func update(screen *ebiten.Image) error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return fmt.Errorf("Window Closed")
	}
	fill(myData.pixelArray, color.Black)
	checkKey()
	drawScreen(screen, myData)
	screen.ReplacePixels(myData.pixelArray)
	return (nil)
}

func main() {
	var err error

	if len(os.Args) < 2 {
		fmt.Printf("usage :\n%s <path_to_map.ini>\n", os.Args[0])
		return
	}
	myData.theMap, err = parser(os.Args[1])
	myData.pixelArray = newPixelArray(width, height)
	if err != nil {
		fmt.Printf("Error : %s\n", err)
	}
	ret := ebiten.Run(update, 400, 400, 1, "Go is Wonderful")
	fmt.Printf("Error : %s\n", ret)
}
