package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Instructions: https://go.dev/tour/methods/25

const w = 10
const h = 5

type Image struct {
	data [][]uint8
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: 0,
		},
		Max: image.Point{
			X: len(img.data),
			Y: len(img.data[0]),
		},
	}
}

func (img Image) At(x, y int) color.Color {
	v := img.data[x][y]
	return color.RGBA{R: v, G: v, B: 255, A: 255}
}

func main() {
	m := Image{}

	m.data = make([][]uint8, h)

	for y := range m.data {
		m.data[y] = make([]uint8, w)

		for x := range m.data[y] {
			m.data[y][x] = uint8(x ^ y)
		}
	}

	pic.ShowImage(m)
}
