package main

import (
	"math"

	"github.com/hajimehoshi/ebiten"
)

func checkKey() {
	var newPos floatPosition

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		newPos = myData.player.pos
		newPos.y = myData.player.pos.y + math.Sin(myData.player.angle*(math.Pi/180))/30
		if checkPos(newPos) {
			myData.player.pos = newPos
		}
		newPos = myData.player.pos
		newPos.x = myData.player.pos.x + math.Cos(myData.player.angle*(math.Pi/180))/30
		if checkPos(newPos) {
			myData.player.pos = newPos
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		newPos = myData.player.pos
		newPos.y = myData.player.pos.y - math.Sin(myData.player.angle*(math.Pi/180))/30
		if checkPos(newPos) {
			myData.player.pos = newPos
		}
		newPos = myData.player.pos
		newPos.x = myData.player.pos.x - math.Cos(myData.player.angle*(math.Pi/180))/30
		if checkPos(newPos) {
			myData.player.pos = newPos
		}
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

func checkPos(pos floatPosition) bool {
	if isThereWall(floatPosToIntPos(pos)) {
		return false
	}
	return true
}
