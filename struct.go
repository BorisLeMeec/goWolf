package main

type position struct {
	x uint32
	y uint32
}

type size position

type floatPosition struct {
	x float64
	y float64
}

type wolfMap struct {
	size  size
	array []uint8
}

type personnage struct {
	angle float64
	pos   position
}

type data struct {
	pixelArray []uint8
	theMap     wolfMap
	player     personnage
	refresh    bool
}
