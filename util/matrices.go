package util

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func AdjacentCells[T any](lines [][]T, x, y int, oob *T) []T {
	res := make([]T, 0, 8)
	// 1 2 3
	// 4 X 5
	// 6 7 8

	appendIfOutOfBounds := func() {
		if oob != nil {
			res = append(res, *oob)
		}
	}

	// 1
	if x > 0 && y > 0 {
		res = append(res, lines[y-1][x-1])
	} else {
		appendIfOutOfBounds()
	}
	// 2
	if y > 0 {
		res = append(res, lines[y-1][x])
	} else {
		appendIfOutOfBounds()
	}
	// 3
	if x < len(lines[y])-1 && y > 0 {
		res = append(res, lines[y-1][x+1])
	} else {
		appendIfOutOfBounds()
	}

	// 4
	if x > 0 {
		res = append(res, lines[y][x-1])
	} else {
		appendIfOutOfBounds()
	}
	// 5
	if x < len(lines[y])-1 {
		res = append(res, lines[y][x+1])
	} else {
		appendIfOutOfBounds()
	}

	// 6
	if x > 0 && y < len(lines)-1 {
		res = append(res, lines[y+1][x-1])
	} else {
		appendIfOutOfBounds()
	}
	// 7
	if y < len(lines)-1 {
		res = append(res, lines[y+1][x])
	} else {
		appendIfOutOfBounds()
	}
	// 8
	if x < len(lines[y])-1 && y < len(lines)-1 {
		res = append(res, lines[y+1][x+1])
	} else {
		appendIfOutOfBounds()
	}
	return res
}

func ConnectingCells[T any](lines [][]T, x, y int, oob *T) []T {
	res := make([]T, 0, 4)
	// x 1 x
	// 2 X 3
	// x 4 x

	appendIfOutOfBounds := func() {
		if oob != nil {
			res = append(res, *oob)
		}
	}

	// 1
	if y > 0 {
		res = append(res, lines[y-1][x])
	} else {
		appendIfOutOfBounds()
	}
	// 2
	if x > 0 {
		res = append(res, lines[y][x-1])
	} else {
		appendIfOutOfBounds()
	}
	// 3
	if x < len(lines[y])-1 {
		res = append(res, lines[y][x+1])
	} else {
		appendIfOutOfBounds()
	}
	// 4
	if y < len(lines)-1 {
		res = append(res, lines[y+1][x])
	} else {
		appendIfOutOfBounds()
	}
	return res
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

type Matrix[T any] struct {
	//   x->
	// y 1 | 0
	// | --+--
	// v 2 | 3
	Quadrants [4]*Quadrant[T]
}

func (m *Matrix[T]) MinMax() (minX, minY, maxX, maxY int) {
	minX = -max(m.Quadrants[1].LenX, m.Quadrants[2].LenX)
	minY = -max(m.Quadrants[1].LenY, m.Quadrants[0].LenY)
	maxX = max(m.Quadrants[0].LenX, m.Quadrants[3].LenX) - 1
	maxY = max(m.Quadrants[2].LenY, m.Quadrants[3].LenY) - 1
	return
}

func NewMatrix[T any]() *Matrix[T] {
	return &Matrix[T]{
		Quadrants: [4]*Quadrant[T]{
			&Quadrant[T]{},
			&Quadrant[T]{},
			&Quadrant[T]{},
			&Quadrant[T]{},
		},
	}
}

func (m Matrix[T]) String() string {
	minX, minY, maxX, maxY := m.MinMax()
	s := strings.Builder{}
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			v, _ := m.Get(x, y)
			s.WriteString(fmt.Sprintf("%v", v))
		}
		s.WriteString("\n")
	}
	return s.String()
}

func getQuadrant(x, y int) (int, int, int) {
	if x < 0 {
		if y < 0 {
			return 1, -x - 1, -y - 1
		}
		return 2, -x - 1, y
	}
	if y < 0 {
		return 0, x, -y - 1
	}
	return 3, x, y
}

func (m *Matrix[T]) Get(x, y int) (T, bool) {
	q, x, y := getQuadrant(x, y)
	return m.Quadrants[q].Get(x, y)
}

func (m *Matrix[T]) Set(x, y int, v T) {
	q, x, y := getQuadrant(x, y)
	m.Quadrants[q].Set(x, y, v)
}

func (m *Matrix[T]) Rows() [][]T {
	minX, minY, maxX, maxY := m.MinMax()
	rows := make([][]T, maxY-minY+1)
	for y := minY; y <= maxY; y++ {
		rows[y-minY] = make([]T, maxX-minX+1)
		for x := minX; x <= maxX; x++ {
			v, _ := m.Get(x, y)
			rows[y-minY][x-minX] = v
		}
	}
	return rows
}

func (m *Matrix[T]) Iterate(f func(x, y int, v T)) {
	minX, minY, maxX, maxY := m.MinMax()
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			v, _ := m.Get(x, y)
			f(x, y, v)
		}
	}
}

func (m *Matrix[T]) ConnectingCells(x, y int) [4]T {
	// x 1 x
	// 2 X 3
	// x 4 x
	var res [4]T
	res[0], _ = m.Get(x, y-1)
	res[1], _ = m.Get(x-1, y)
	res[2], _ = m.Get(x+1, y)
	res[3], _ = m.Get(x, y+1)
	return res
}
