package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten"
)

// PixelArray represents an array of pixel.
type PixelArray struct {
	pixels []uint8
	size   size
	scale  size
	rotate float64
}

func getPixelArray(img *ebiten.Image) (PixelArray, error) {
	imageHeight, imageWidth := img.Size()
	var out PixelArray

	out.size.x = uint32(imageWidth)
	out.size.y = uint32(imageHeight)
	for y := 0; y < imageHeight; y += 4 {
		for x := 0; x < imageWidth; x += 4 {
			r, g, b, a := img.At(y, x).RGBA()
			out.pixels = append(out.pixels, uint8(r>>8), uint8(g>>8), uint8(b>>8), uint8(a>>8))
		}
	}
	return out, nil
}

// SetPixel set a pixel at potition 'pos' of color 'color'
func (pix *PixelArray) SetPixel(pos position, color color.Color) {
	if pos.x < 0 || pos.y < 0 || pos.x > pix.size.x-1 || pos.y > pix.size.y-1 {
		return
	}
	index := 4 * (pos.y*pix.size.x + pos.x)
	r, g, b, a := color.RGBA()
	pix.setColorAt(index, uint8(r>>8), uint8(g>>8), uint8(b>>8), uint8(a>>8))
}

// NewPixelArray return a pixelArray of size (width, height)
func NewPixelArray(width, height uint32) PixelArray {
	var out PixelArray

	out.size.x = width
	out.size.y = height
	out.scale = size{1, 1}
	for y := uint32(0); y < height*4; y += 4 {
		for x := uint32(0); x < width*4; x += 4 {
			r, g, b, a := color.Black.RGBA()
			out.pixels = append(out.pixels, uint8(r>>8), uint8(g>>8), uint8(b>>8), uint8(a>>8))
		}
	}
	return out
}

// SetScale is used to set the scale of 'pix'
func (pix *PixelArray) SetScale(x, y uint32) {
	if x < 0 || y < 0 {
		return
	}
	pix.scale = size{x, y}
}

// SetRotate is used to set the rotation of 'pix'
func (pix *PixelArray) SetRotate(r float64) {
	if r < 0 {
		return
	}
	pix.rotate = r
}

// GetRotate return the current value of 'pix''srotate'
func (pix *PixelArray) GetRotate() float64 {
	return pix.rotate
}

// Blit copy array 'src' over pix from 'posStart' to 'posStart' + 'size'
func (pix *PixelArray) Blit(src PixelArray, posStart position, size size) {
	var rSrc, gSrc, bSrc, aSrc, rBlit, gBlit, bBlit, aBlit, rDest, gDest, bDest, aDest uint8
	var posToBlit, posToGet position
	var indexBlit, indexGet uint32

	if posStart.x < 0 || posStart.x > pix.size.x || posStart.y < 0 || posStart.y > pix.size.y {
		return
	}
	if size.x < 0 || size.y < 0 {
		return
	}
	cos := math.Cos(src.rotate * (math.Pi / 180))
	sin := math.Sin(src.rotate * (math.Pi / 180))
	centerX := src.size.x / 2
	centerY := src.size.y / 2

	for posToGet.y = 0; posToGet.y < src.size.y && posToGet.y < size.y; posToGet.y++ {
		for posToGet.x = 0; posToGet.x < src.size.x && posToGet.x < size.x; posToGet.x++ {
			indexGet = 4 * (posToGet.y*src.size.x + posToGet.x)
			rSrc, gSrc, bSrc, aSrc = src.getColorAt(indexGet)
			for i := uint32(0); i < src.scale.x; i++ {
				for j := uint32(0); j < src.scale.y; j++ {
					m := src.scale.x*(posToGet.x+posStart.x) + i - centerX
					n := src.scale.y*(posToGet.y+posStart.y) + j - centerY
					posToBlit.x = uint32((float64(m)*cos + float64(n)*sin)) + centerX
					posToBlit.y = uint32((float64(n)*cos - float64(m)*sin)) + centerY

					if posToBlit.x > pix.size.x-1 || posToBlit.y > pix.size.y-1 {
						continue
					}
					indexBlit = 4 * (posToBlit.y*pix.size.x + posToBlit.x)
					rDest, gDest, bDest, aDest = pix.getColorAt(indexBlit)
					rBlit = uint8(float32(rSrc)*(float32(aSrc)/255.0) + float32(rDest)*(1.0-(float32(aSrc)/255.0)))
					gBlit = uint8(float32(gSrc)*(float32(aSrc)/255.0) + float32(gDest)*(1.0-(float32(aSrc)/255.0)))
					bBlit = uint8(float32(bSrc)*(float32(aSrc)/255.0) + float32(bDest)*(1.0-(float32(aSrc)/255.0)))
					aBlit = aDest
					pix.setColorAt(indexBlit, rBlit, gBlit, bBlit, aBlit)
				}
			}
		}
	}
}

// Fill fill every pix'pixels with color 'color'
func (pix *PixelArray) Fill(color color.Color) {
	_r, _g, _b, _a := color.RGBA()

	for x := uint32(0); x < pix.size.x*pix.size.y*4; x += 4 {
		pix.setColorAt(x, uint8(_r>>8), uint8(_g>>8), uint8(_b>>8), uint8(_a>>8))
	}
}
