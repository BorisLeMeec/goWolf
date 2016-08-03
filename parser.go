package main

import (
	"fmt"
	"strconv"
	"strings"

	ini "gopkg.in/ini.v1"
)

func parser(filename string) (wolfMap, error) {
	var out wolfMap
	ini, err := ini.Load(filename)

	if err != nil {
		return out, err
	}
	width, err := ini.Section("").GetKey("width")
	if err != nil {
		return out, err
	}
	height, err := ini.Section("").GetKey("height")
	if err != nil {
		return out, err
	}
	carte, err := ini.Section("").GetKey("data")
	if err != nil {
		return out, err
	}
	carteFormat := strings.Replace(carte.Value(), ",", "", -1)
	heightInt, _ := strconv.Atoi(height.Value())
	widthInt, _ := strconv.Atoi(width.Value())
	if len(carteFormat) != heightInt*widthInt {
		return out, fmt.Errorf("Map format doesn't fit")
	}
	out.size.x = uint32(widthInt)
	out.size.y = uint32(heightInt)
	out.array = []byte(carteFormat)
	return out, nil
}
