package util

import (
	"math"

	"golang.org/x/exp/constraints"
)

func ShoelaceArea[T int | constraints.Integer](path [][2]T) float64 {
	path = append(path, path[0])
	area := .0
	prevX, prevY := path[0][0], path[0][1]
	for i, point := range path {
		if i == 0 {
			continue
		}
		x, y := point[0], point[1]
		area += float64(prevX)*float64(y) - float64(prevY)*float64(x)
		prevX, prevY = x, y
	}
	return math.Abs(area / 2)
}
