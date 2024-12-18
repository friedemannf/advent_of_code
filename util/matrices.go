package util

import (
	"fmt"
	"iter"
	"slices"

	"github.com/fatih/color"
)

type Matrix[T comparable] [][]T

// Width returns the width of the matrix (Y)
// Returns 0 if the matrix is empty
func (m Matrix[T]) Width() int {
	return len(m)
}

// Height returns the height of the matrix (X)
// Panics if the matrix is malformed, i.e. there are lines of different length
func (m Matrix[T]) Height() int {
	width := 0
	for _, line := range m {
		if width != 0 && len(line) != width {
			panic("Invalid matrix")
		}
		width = len(line)
	}
	return width
}

func (matrix Matrix[T]) Copy() Matrix[T] {
	m := make([][]T, len(matrix))
	for i, l := range matrix {
		m[i] = make([]T, len(l))
		copy(m[i], l)
	}
	return m
}

func MatrixFromLines(lines []string) Matrix[rune] {
	m := Matrix[rune]{}
	for _, line := range lines {
		m = append(m, []rune(line))
	}
	return m
}

func MakeMatrix[T comparable](height, width int) Matrix[T] {
	m := make([][]T, height)
	for i := range m {
		m[i] = make([]T, width)
	}
	return m
}

func (matrix Matrix[T]) AdjacentCells(coord Coordinate, oob *T) []T {
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
	if coord.X > 0 && coord.Y > 0 {
		res = append(res, matrix[coord.Y-1][coord.X-1])
	} else {
		appendIfOutOfBounds()
	}
	// 2
	if coord.Y > 0 {
		res = append(res, matrix[coord.Y-1][coord.X])
	} else {
		appendIfOutOfBounds()
	}
	// 3
	if coord.X < len(matrix[coord.Y])-1 && coord.Y > 0 {
		res = append(res, matrix[coord.Y-1][coord.X+1])
	} else {
		appendIfOutOfBounds()
	}

	// 4
	if coord.X > 0 {
		res = append(res, matrix[coord.Y][coord.X-1])
	} else {
		appendIfOutOfBounds()
	}
	// 5
	if coord.X < len(matrix[coord.Y])-1 {
		res = append(res, matrix[coord.Y][coord.X+1])
	} else {
		appendIfOutOfBounds()
	}

	// 6
	if coord.X > 0 && coord.Y < len(matrix)-1 {
		res = append(res, matrix[coord.Y+1][coord.X-1])
	} else {
		appendIfOutOfBounds()
	}
	// 7
	if coord.Y < len(matrix)-1 {
		res = append(res, matrix[coord.Y+1][coord.X])
	} else {
		appendIfOutOfBounds()
	}
	// 8
	if coord.X < len(matrix[coord.Y])-1 && coord.Y < len(matrix)-1 {
		res = append(res, matrix[coord.Y+1][coord.X+1])
	} else {
		appendIfOutOfBounds()
	}
	return res
}

func (matrix Matrix[T]) ConnectingCells(coord Coordinate, oob *T) []T {
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
	if coord.Y > 0 {
		res = append(res, matrix[coord.Y-1][coord.X])
	} else {
		appendIfOutOfBounds()
	}
	// 2
	if coord.X > 0 {
		res = append(res, matrix[coord.Y][coord.X-1])
	} else {
		appendIfOutOfBounds()
	}
	// 3
	if coord.X < len(matrix[coord.Y])-1 {
		res = append(res, matrix[coord.Y][coord.X+1])
	} else {
		appendIfOutOfBounds()
	}
	// 4
	if coord.Y < len(matrix)-1 {
		res = append(res, matrix[coord.Y+1][coord.X])
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

func (matrix Matrix[T]) Get(point Coordinate) (T, bool) {
	if point.Y < 0 || point.X < 0 {
		return Null[T](), false
	}
	if len(matrix) <= point.Y {
		return Null[T](), false
	}
	if len(matrix[0]) <= point.X {
		return Null[T](), false
	}
	return matrix[point.Y][point.X], true
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

func (matrix Matrix[T]) Iterate() iter.Seq2[Coordinate, T] {
	return func(yield func(Coordinate, T) bool) {
		coordinate := Coordinate{}
		for coordinate.Y = 0; coordinate.Y < len(matrix); coordinate.Y++ {
			for coordinate.X = 0; coordinate.X < len(matrix); coordinate.X++ {
				v, _ := matrix.Get(coordinate)
				if !yield(coordinate, v) {
					return
				}
			}
		}
	}
}

func (matrix Matrix[T]) FloodFill(coordinate Coordinate) []Coordinate {
	m := matrix.Copy()
	return m.floodFill(coordinate)
}

func (matrix Matrix[T]) floodFill(coordinate Coordinate) []Coordinate {
	val, ok := matrix.Get(coordinate)
	if !ok {
		return nil
	}
	if val == Null[T]() {
		panic(fmt.Sprintf("called with zero value: X=%v Y=%v val=%v", coordinate.X, coordinate.Y, val))
	}
	c := []Coordinate{coordinate}
	matrix.Set(coordinate, Null[T]())
	// Up
	if v, ok := matrix.Get(coordinate.Add(VecUp)); ok && v == val {
		c = append(c, matrix.floodFill(coordinate.Add(VecUp))...)
	}
	// Left
	if v, ok := matrix.Get(coordinate.Add(VecLeft)); ok && v == val {
		c = append(c, matrix.floodFill(coordinate.Add(VecLeft))...)
	}
	// Right
	if v, ok := matrix.Get(coordinate.Add(VecRight)); ok && v == val {
		c = append(c, matrix.floodFill(coordinate.Add(VecRight))...)
	}
	// Down
	if v, ok := matrix.Get(coordinate.Add(VecDown)); ok && v == val {
		c = append(c, matrix.floodFill(coordinate.Add(VecDown))...)
	}
	return c
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

var (
	VecUp    Vector = Vector{X: 0, Y: -1}
	VecDown  Vector = Vector{X: 0, Y: 1}
	VecLeft  Vector = Vector{X: -1, Y: 0}
	VecRight Vector = Vector{X: 1, Y: 0}
)

func (c Coordinate) DistanceTo(to Coordinate) Vector {
	return Vector{
		X: to.X - c.X,
		Y: to.Y - c.Y,
	}
}

type MChar int32

func (m MChar) String() string {
	switch m {
	case 0:
		return color.HiBlackString(".")
	case '#':
		return color.HiRedString("#")
	case 'X':
		return color.HiGreenString("X")
	}
	return string(m)
}

type Quadrant[T any] struct {
	LenX, LenY int
	Rows       [][]T
}

func (q *Quadrant[T]) Get(x, y int) (T, bool) {
	var v T
	if q.LenY <= y || q.LenX <= x {
		return v, false
	}
	if len(q.Rows) <= y || len(q.Rows[y]) <= x {
		return v, true
	}
	return q.Rows[y][x], true
}

func (q *Quadrant[T]) Set(x, y int, v T) {
	if len(q.Rows) <= y {
		for i := len(q.Rows); i <= y; i++ {
			q.Rows = append(q.Rows, nil)
		}
	}
	if len(q.Rows[y]) <= x {
		q.Rows[y] = append(q.Rows[y], make([]T, x-len(q.Rows[y])+1)...)
	}
	q.Rows[y][x] = v
	q.LenX = max(q.LenX, x+1)
	q.LenY = max(q.LenY, y+1)
}
