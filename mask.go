package main

import (
	"fmt"
	"mime"
)

// Mask is used to hide or show a part of a PixelArray
// By default it show everything
type Mask struct {
	value  []bool
	isUsed bool
}

func (p *PixelArray) newMask() {
	p.mask.isUsed = true
	p.mask.value = make([]bool, p.size.x*p.size.y)
	for i := 0; i < len(p.mask.value); i++ {
		p.mask.value[i] = true
	}
}

// ApplyMask load the 'file' and apply it  to pixArray p
func (p *PixelArray) ApplyMask(file string) {
	mime := mime.TypeByExtension(file)
	fmt.Println(mime)
}
