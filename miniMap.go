package main

import "image/color"

func createMiniMap() miniMap {
	var out miniMap

	out.pix = newPixelArray(10, 10)
	out.pix.scale = newSize(5, 5)
	out.zoom = 1
	out.posStart = newPosition(0, 0)
	return out
}

func drawMiniMap(myData data) {
	var index uint32
	var posInMap position
	var myColor color.Color

	for posInMap.y = uint32(0); posInMap.y < myData.theMap.size.y; posInMap.y++ {
		for posInMap.x = uint32(0); posInMap.x < myData.theMap.size.x; posInMap.x++ {
			index = posInMap.y*myData.theMap.size.x + posInMap.x
			switch myData.theMap.array[index] {
			case '1':
				myColor = color.White
			default:
				myColor = color.Black
			}
			setPixel(myData.miniMap.pix, posInMap, myColor)
		}
	}
}
