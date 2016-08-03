package main

func getColorAt(pix pixelArray, index uint32) (r, g, b, a uint8) {
	r = pix.pixels[index+0]
	g = pix.pixels[index+1]
	b = pix.pixels[index+2]
	a = pix.pixels[index+3]
	return
}

func setColorAt(pix pixelArray, index uint32, r, g, b, a uint8) {
	pix.pixels[index+0] = r
	pix.pixels[index+1] = g
	pix.pixels[index+2] = b
	pix.pixels[index+3] = a
}
