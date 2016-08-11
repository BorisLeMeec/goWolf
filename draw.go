package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

func drawScreen(img *ebiten.Image) {
	drawWalls()
	drawMiniMap()
	pos := newPosition(int(myData.player.pos.x*15), int(myData.player.pos.y*15))
	setPixel(myData.pix, pos, color.White)
	blit(myData.pix, myData.miniMap.pix, myData.miniMap.posStart, myData.miniMap.pix.size)
}
