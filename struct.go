package main

type pixelArray struct {
	size   size
	pixels []uint8
	scale  size
}

type position struct {
	x uint32
	y uint32
}

func (in *floatPosition) toIntPos() position {
	var out position

	out.x = uint32(in.x)
	out.y = uint32(in.y)
	return out
}

type size position

type floatPosition struct {
	x float64
	y float64
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
	pix      pixelArray
	zoom     float32
}

type data struct {
	pix     pixelArray
	miniMap miniMap
	theMap  wolfMap
	player  personnage
	refresh bool
}
