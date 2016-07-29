package main

import (
	"fmt"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten"
)

var myData data

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
	pixels, _ := getPixelArray(screen)

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return fmt.Errorf("Window Closed")
	}
	screen.Fill(color.Black)
	checkKey()
	setPixel(screen, pixels, myData.player.pos, color.White)
	screen.ReplacePixels(pixels)
	return (nil)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage :\n%s <path_to_map.ini>\n", os.Args[0])
		return
	}
	_, err := parser(os.Args[1])
	fmt.Printf("Error : %s\n", err)
	ret := ebiten.Run(update, 400, 400, 1, "Go is Wonderful")
	fmt.Printf("Error : %s\n", ret)
}
