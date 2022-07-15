package main

import (
	"golang.org/x/tour/pic"
)

// Instructions: https://go.dev/tour/moretypes/18

func Pic(dx, dy int) (img [][]uint8) {
	for x := 0; x < dx; x++ {
		img = append(img, make([]uint8, dy))

		for y := 0; y < dy; y++ {
			img[x][y] = uint8(x ^ y)
		}
	}

	return
}

func main() {
	pic.Show(Pic)
}
