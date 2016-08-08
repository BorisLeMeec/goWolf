package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

func setPixel(pix pixelArray, pos position, color color.Color) {
	if pos.x < 0 || pos.y < 0 || pos.x > pix.size.x-1 || pos.y > pix.size.y-1 {
		return
	}
	index := 4 * (pos.y*pix.size.x + pos.x)
	r, g, b, a := color.RGBA()
	setColorAt(pix, index, uint8(r>>8), uint8(g>>8), uint8(b>>8), uint8(a>>8))
}

func getPixelArray(img *ebiten.Image) (pixelArray, error) {
	imageHeight, imageWidth := img.Size()
	var out pixelArray

	for y := 0; y < imageHeight; y += 4 {
		for x := 0; x < imageWidth; x += 4 {
			r, g, b, a := img.At(y, x).RGBA()
			out.pixels = append(out.pixels, uint8(r>>8), uint8(g>>8), uint8(b>>8), uint8(a>>8))
		}
	}
	return out, nil
}

func newPixelArray(width, height uint32) pixelArray {
	var out pixelArray

	out.size.x = width
	out.size.y = height
	for y := uint32(0); y < height*4; y += 4 {
		for x := uint32(0); x < width*4; x += 4 {
			r, g, b, a := color.Black.RGBA()
			out.pixels = append(out.pixels, uint8(r>>8), uint8(g>>8), uint8(b>>8), uint8(a>>8))
		}
	}
	return out
}

func blit(dest, src pixelArray, posStart position, size size) {
	var rSrc, gSrc, bSrc, aSrc, rBlit, gBlit, bBlit, aBlit, rDest, gDest, bDest, aDest uint8
	var posToBlit, posToGet position
	var indexBlit, indexGet uint32

	if posStart.x < 0 || posStart.x > dest.size.x || posStart.y < 0 || posStart.y > dest.size.y {
		return
	}
	if size.x < 0 || size.y < 0 {
		return
	}
	for posToGet.y = 0; posToGet.y < src.size.y && posToGet.y < size.y; posToGet.y++ {
		for posToGet.x = 0; posToGet.x < src.size.x && posToGet.x < size.x; posToGet.x++ {
			indexGet = 4 * (posToGet.y*src.size.x + posToGet.x)
			rSrc, gSrc, bSrc, aSrc = getColorAt(src, indexGet)
			for i := uint32(0); i < src.scale.x; i++ {
				for j := uint32(0); j < src.scale.y; j++ {
					posToBlit.x = src.scale.x*(posToGet.x+posStart.x) + j
					posToBlit.y = src.scale.y*(posToGet.y+posStart.y) + i
					if posToBlit.x > dest.size.x-1 || posToBlit.y > dest.size.y-1 {
						continue
					}
					indexBlit = 4 * (posToBlit.y*dest.size.x + posToBlit.x)
					rDest, gDest, bDest, aDest = getColorAt(dest, indexBlit)
					rBlit = uint8(float32(rSrc)*(float32(aSrc)/255.0) + float32(rDest)*(1.0-(float32(aSrc)/255.0)))
					gBlit = uint8(float32(gSrc)*(float32(aSrc)/255.0) + float32(gDest)*(1.0-(float32(aSrc)/255.0)))
					bBlit = uint8(float32(bSrc)*(float32(aSrc)/255.0) + float32(bDest)*(1.0-(float32(aSrc)/255.0)))
					aBlit = aDest
					setColorAt(dest, indexBlit, rBlit, gBlit, bBlit, aBlit)
				}
			}
		}
	}
}

func fill(pix pixelArray, color color.Color) {
	_r, _g, _b, _a := color.RGBA()

	for x := uint32(0); x < pix.size.x*pix.size.y*4; x += 4 {
		setColorAt(pix, x, uint8(_r>>8), uint8(_g>>8), uint8(_b>>8), uint8(_a>>8))
	}
}
