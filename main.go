package main

import (
	"fmt"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten"
)

var myData data
var height = uint32(400)
var width = uint32(400)
var angleBetweenRay = 60 / float64(width)

func update(screen *ebiten.Image) error {
	myData.miniMap.pix.rotate++
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return fmt.Errorf("Window Closed")
	}
	myData.pix.Fill(color.Black)
	checkKey()
	drawScreen(screen)
	screen.ReplacePixels(myData.pix.pixels)
	return (nil)
}

func main() {
	var err error

	if len(os.Args) < 2 {
		fmt.Printf("usage :\n%s <path_to_map.ini>\n", os.Args[0])
		return
	}
	myData.theMap, err = parser(os.Args[1])
	myData.pix = NewPixelArray(width, height)
	myData.miniMap = createMiniMap()
	myData.player.pos = floatPosition{2, 2}
	if err != nil {
		fmt.Printf("Error : %s\n", err)
	}
	ret := ebiten.Run(update, int(width), int(height), 1, "Go is Wonderful")
	fmt.Printf("Error : %s\n", ret)
}
