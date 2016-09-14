package main

import "github.com/hajimehoshi/ebiten"

func drawScreen(img *ebiten.Image) {
	drawWalls()
	drawMiniMap()
	myData.pix.Blit(*myData.miniMap.pix, myData.miniMap.posStart, myData.miniMap.pix.size)
}
