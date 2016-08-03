package main

import (
	"math"

	"github.com/hajimehoshi/ebiten"
)

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
		myData.player.angle -= 3
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		myData.player.angle += 3
	}
	if myData.player.angle > 360 {
		myData.player.angle = 0
	}
	if myData.player.angle < 0 {
		myData.player.angle = 360
	}
}
