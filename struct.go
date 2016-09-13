package main

type position struct {
	x int
	y int
}

type size position

func (p *size) maxSize() int {
	return p.x*p.y - 1
}

type floatPosition struct {
	x float64
	y float64
}

func (in *floatPosition) toIntPos() position {
	var out position

	out.x = int(in.x)
	out.y = int(in.y)
	return out
}

type vect floatPosition

type ray struct {
	pos    position
	dist   float32
	angle  float64
	vecDir vect
}

type wolfMap struct {
	size  size
	array []byte
}

type personnage struct {
	angle float64
	pos   floatPosition
}

type miniMap struct {
	posStart position
	pix      PixelArray
	zoom     float32
}

type data struct {
	pix     PixelArray
	miniMap miniMap
	theMap  wolfMap
	player  personnage
	refresh bool
}
