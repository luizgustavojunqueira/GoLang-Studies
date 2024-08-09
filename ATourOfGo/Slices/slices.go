package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	var s = make([][]uint8, dy)

	for i := range dy {
		for range dx {
			s[i] = append(s[i], 0)
		}
	}

	return s
}

func main() {
	pic.Show(Pic)
}
