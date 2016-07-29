package main

import (
	"fmt"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten"
)

var myData data
var height = 400
var width = 400

func checkKey() {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		myData.player.pos.y--
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		myData.player.pos.y++
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		myData.player.pos.x--
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		myData.player.pos.x++
	}
}

func update(screen *ebiten.Image) error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return fmt.Errorf("Window Closed")
	}
	fill(myData.pixelArray, color.Black)
	checkKey()
	drawWall(screen, myData.pixelArray, 100, 200)
	setPixel(screen, myData.pixelArray, myData.player.pos, color.White)
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
