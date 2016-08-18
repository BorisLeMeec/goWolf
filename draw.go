package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

func drawScreen(img *ebiten.Image) {
	drawWalls()
	drawMiniMap()
	pos := position{uint32(myData.player.pos.x * 15), uint32(myData.player.pos.y * 15)}
	myData.pix.SetPixel(pos, color.White)
	myData.pix.Blit(myData.miniMap.pix, myData.miniMap.posStart, myData.miniMap.pix.size)
}
