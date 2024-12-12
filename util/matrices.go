package util

import (
	"fmt"
	"slices"

	"github.com/fatih/color"
)

type Matrix[T comparable] [][]T

func (matrix Matrix[T]) AdjacentCells(x, y int, oob *T) []T {
	res := make([]T, 0, 8)
	// 0 1 2
	// 3 X 4
	// 5 6 7

	appendIfOutOfBounds := func() {
		if oob != nil {
			res = append(res, *oob)
		}
	}

	// 1
	if x > 0 && y > 0 {
		res = append(res, matrix[y-1][x-1])
	} else {
		appendIfOutOfBounds()
	}
	// 2
	if y > 0 {
		res = append(res, matrix[y-1][x])
	} else {
		appendIfOutOfBounds()
	}
	// 3
	if x < len(matrix[y])-1 && y > 0 {
		res = append(res, matrix[y-1][x+1])
	} else {
		appendIfOutOfBounds()
	}

	// 4
	if x > 0 {
		res = append(res, matrix[y][x-1])
	} else {
		appendIfOutOfBounds()
	}
	// 5
	if x < len(matrix[y])-1 {
		res = append(res, matrix[y][x+1])
	} else {
		appendIfOutOfBounds()
	}

	// 6
	if x > 0 && y < len(matrix)-1 {
		res = append(res, matrix[y+1][x-1])
	} else {
		appendIfOutOfBounds()
	}
	// 7
	if y < len(matrix)-1 {
		res = append(res, matrix[y+1][x])
	} else {
		appendIfOutOfBounds()
	}
	// 8
	if x < len(matrix[y])-1 && y < len(matrix)-1 {
		res = append(res, matrix[y+1][x+1])
	} else {
		appendIfOutOfBounds()
	}
	return res
}

func (matrix Matrix[T]) ConnectingCells(x, y int, oob *T) []T {
	res := make([]T, 0, 4)
	// x 0 x
	// 1 X 2
	// x 3 x

	appendIfOutOfBounds := func() {
		if oob != nil {
			res = append(res, *oob)
		}
	}

	// 1
	if y > 0 {
		res = append(res, matrix[y-1][x])
	} else {
		appendIfOutOfBounds()
	}
	// 2
	if x > 0 {
		res = append(res, matrix[y][x-1])
	} else {
		appendIfOutOfBounds()
	}
	// 3
	if x < len(matrix[y])-1 {
		res = append(res, matrix[y][x+1])
	} else {
		appendIfOutOfBounds()
	}
	// 4
	if y < len(matrix)-1 {
		res = append(res, matrix[y+1][x])
	} else {
		appendIfOutOfBounds()
	}
	return res
}

func (matrix Matrix[T]) Contains(point Coordinate) bool {
	return len(matrix) > 0 &&
		point.Y >= 0 && point.Y < len(matrix) &&
		point.X >= 0 && point.X < len(matrix[0])
}

func (matrix Matrix[T]) Set(point Coordinate, to T) bool {
	if !matrix.Contains(point) {
		return false
	}
	matrix[point.Y][point.X] = to
	return true
}

func (m Matrix[T]) ColorPrint(colors map[T]color.Attribute) {
	for _, line := range m {
		for _, char := range line {
			switch r := any(char).(type) {
			case rune:
				// Special case for runes
				if c, ok := colors[char]; ok {
					color.New(c).Print(string(r))
				} else {
					nul := new(T)
					if dflt, ok := colors[*nul]; ok {
						color.New(dflt).Print(string(r))
					} else {
						fmt.Print(string(r))
					}
				}
			default:
				if c, ok := colors[char]; ok {
					color.New(c).Print(char)
				} else {
					nul := new(T)
					if dflt, ok := colors[*nul]; ok {
						color.New(dflt).Print(r)
					} else {
						fmt.Print(r)
					}
				}
			}
		}
		fmt.Println()
	}
}

type Modifier uint8

const (
	// 1 →→→
	// 2 →→→
	// 3 →→→
	Horizontal Modifier = 2 << iota
	// 3456
	// 2 ↘↘↘↘
	// 1 ↘↘↘↘
	//   ↘↘↘↘
	DiagonalLR
	//   6543
	// ↙↙↙↙ 2
	// ↙↙↙↙ 1
	// ↙↙↙↙
	DiagonalRL
	// 123
	// ↓↓↓
	// ↓↓↓
	// ↓↓↓
	Vertical
	Forward Modifier = 0
	Reverse Modifier = 1
)

func (matrix Matrix[T]) Transform(dir Modifier) (out [][]T) {
	reverseIf := func(l []T) {
		if dir&Reverse > 0 {
			slices.Reverse(l)
		}
	}
	if dir&Horizontal > 0 {
		out = make([][]T, len(matrix))
		for i, line := range matrix {
			out[i] = make([]T, len(line))
			for i2, t := range line {
				out[i][i2] = t
			}
			reverseIf(out[i])
		}
	} else if dir&Vertical > 0 {
		x, y := len(matrix[0]), len(matrix)
		out = make([][]T, x)
		for i := 0; i < x; i++ {
			out[i] = make([]T, y)
			for ii := 0; ii < y; ii++ {
				out[i][ii] = matrix[ii][i]
			}
			reverseIf(out[i])
		}
	} else if dir&DiagonalLR > 0 {
		out = make([][]T, 0, len(matrix)+len(matrix[0])-1)
		for i := len(matrix) - 1; i >= 0; i-- {
			temp := make([]T, 0)
			for ii := 0; ii < len(matrix[0]) && (i+ii < len(matrix)); ii++ {
				temp = append(temp, matrix[i+ii][ii])
			}
			reverseIf(temp)
			out = append(out, temp)
		}
		for i := 1; i < len(matrix[0]); i++ {
			temp := make([]T, 0)
			for ii := 0; ii < len(matrix) && (i+ii < len(matrix[0])); ii++ {
				temp = append(temp, matrix[ii][i+ii])
			}
			reverseIf(temp)
			out = append(out, temp)
		}
	} else if dir&DiagonalRL > 0 {
		out = make([][]T, 0, len(matrix)+len(matrix[0])-1)
		for i := len(matrix) - 1; i >= 0; i-- {
			temp := make([]T, 0)
			for ii := 0; ii < len(matrix[0]) && (i+ii < len(matrix)); ii++ {
				temp = append(temp, matrix[i+ii][len(matrix[0])-1-ii])
			}
			reverseIf(temp)
			out = append(out, temp)
		}
		for i := len(matrix[0]) - 2; i >= 0; i-- {
			temp := make([]T, 0)
			for ii := 0; ii < len(matrix) && (i-ii >= 0); ii++ {
				temp = append(temp, matrix[ii][i-ii])
			}
			reverseIf(temp)
			out = append(out, temp)
		}
	} else {
		panic("Unknown modifier")
	}
	return out
}

type Coordinate struct {
	X, Y int
}

func (c Coordinate) Equal(c2 Coordinate) bool {
	return c.X == c2.X && c.Y == c2.Y
}

func (c Coordinate) Copy() Coordinate {
	return Coordinate{
		X: c.X,
		Y: c.Y,
	}
}

func (c Coordinate) Add(vec Vector) Coordinate {
	return Coordinate{
		X: c.X + vec.X,
		Y: c.Y + vec.Y,
	}
}

func (c Coordinate) Sub(vec Vector) Coordinate {
	return Coordinate{
		X: c.X - vec.X,
		Y: c.Y - vec.Y,
	}
}

type Vector Coordinate

func (c Coordinate) DistanceTo(to Coordinate) Vector {
	return Vector{
		X: to.X - c.X,
		Y: to.Y - c.Y,
	}
}
