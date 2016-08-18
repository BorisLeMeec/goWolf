package main

import "image/color"

func createMiniMap() miniMap {
	var out miniMap

	out.pix = NewPixelArray(10, 10)
	out.pix.scale = size{15, 15}
	out.zoom = 1
	out.posStart = position{0, 0}
	out.pix.SetRotate(45)
	return out
}

func drawMiniMap() {
	var index uint32
	var posPlayer, posInMap position
	var myColor color.Color

	posPlayer = myData.player.pos.toIntPos()
	for posInMap.y = uint32(0); posInMap.y < myData.theMap.size.y; posInMap.y++ {
		for posInMap.x = uint32(0); posInMap.x < myData.theMap.size.x; posInMap.x++ {
			index = posInMap.y*myData.theMap.size.x + posInMap.x
			switch myData.theMap.array[index] {
			case '1':
				myColor = color.RGBA{255, 255, 255, 255}
			default:
				myColor = color.RGBA{0, 0, 0, 255}
			}
			if posInMap.x == posPlayer.x && posInMap.y == posPlayer.y {
				myColor = color.RGBA{255, 0, 0, 255}
			}
			myData.miniMap.pix.SetPixel(posInMap, myColor)
		}
	}
	myData.miniMap.pix.changeOpacity(0.6)
}
