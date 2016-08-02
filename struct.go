package main

func floatPosToIntPos(in floatPosition) position {
	var out position

	out.x = uint32(in.x)
	out.y = uint32(in.y)
	return out
}

type pixelArray struct {
	size   size
	pixels []uint8
}

type position struct {
	x uint32
	y uint32
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
	size size
	pix  pixelArray
}

type personnage struct {
	angle float64
	pos   floatPosition
}

type data struct {
	pix     pixelArray
	theMap  wolfMap
	player  personnage
	refresh bool
}
