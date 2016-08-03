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
	pix.pixels[index+0] = uint8(r)
	pix.pixels[index+1] = uint8(g)
	pix.pixels[index+2] = uint8(b)
	pix.pixels[index+3] = uint8(a)
}

func getPixelArray(img *ebiten.Image) (pixelArray, error) {
	var out pixelArray

	imageHeight, imageWidth := img.Size()

	for y := 0; y < imageHeight; y += 4 {
		for x := 0; x < imageWidth; x += 4 {
			r, g, b, a := img.At(y, x).RGBA()
			out.pixels = append(out.pixels, uint8(r), uint8(g), uint8(b), uint8(a))
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
			out.pixels = append(out.pixels, uint8(r), uint8(g), uint8(b), uint8(a))
		}
	}
	return out
}

func blit(dest, src pixelArray, posStart position, size size) {
	var posToBlit, posToGet position
	var indexBlit, indexGet uint32
	var rSrc, gSrc, bSrc, aSrc, rBlit, gBlit, bBlit, aBlit uint8

	if posStart.x < 0 || posStart.x > dest.size.x || posStart.y < 0 || posStart.y > dest.size.y {
		return
	}
	if size.x < 0 || size.y < 0 {
		return
	}
	for posToGet.y = 0; posToGet.y < src.size.y && posToGet.y < size.y; posToGet.y++ {
		for posToGet.x = 0; posToGet.x < src.size.x && posToGet.x < size.x; posToGet.x++ {
			indexGet = 4 * (posToGet.y*src.size.x + posToGet.x)
			posToBlit.y = posToGet.y + posStart.y
			posToBlit.x = posToGet.x + posStart.x
			if posToBlit.x > dest.size.x-1 || posToBlit.y > dest.size.y-1 {
				continue
			}
			indexBlit = 4 * (posToBlit.y*dest.size.x + posToBlit.x)
			rSrc = src.pixels[indexGet+0]
			gSrc = src.pixels[indexGet+1]
			bSrc = src.pixels[indexGet+2]
			aSrc = src.pixels[indexGet+3]
			// rDest = src.pixels[indexBlit+0]
			// gDest = src.pixels[indexBlit+1]
			// bDest = src.pixels[indexBlit+2]
			// aDest = src.pixels[indexBlit+3]
			rBlit = rSrc
			gBlit = gSrc
			bBlit = bSrc
			aBlit = aSrc
			dest.pixels[indexBlit+0] = rBlit
			dest.pixels[indexBlit+1] = gBlit
			dest.pixels[indexBlit+2] = bBlit
			dest.pixels[indexBlit+3] = aBlit
		}
	}
}

func fill(pix pixelArray, color color.Color) {
	_r, _g, _b, _a := color.RGBA()
	r := uint8(_r)
	g := uint8(_g)
	b := uint8(_b)
	a := uint8(_a)
	for x := uint32(0); x < pix.size.x*pix.size.y*4; x += 4 {
		pix.pixels[x+0] = r
		pix.pixels[x+1] = g
		pix.pixels[x+2] = b
		pix.pixels[x+3] = a
	}
}
