package day10

import (
	"fmt"

	"github.com/fatih/color"

	"github.com/friedemannf/advent_of_code/day"
	"github.com/friedemannf/advent_of_code/util"
)

func init() {
	day.Register(2023, 10, day.Day{
		Solution1: []day.Solution{solution1},
		Solution2: []day.Solution{solution2, solution22},
	})
}

type Char struct {
	C         int32
	Visited   bool
	Contained bool
}

func (c Char) String() string {
	if c.Contained {
		return color.BlueString(string(c.C))
	}
	greenIfVisited := func(s string) string {
		if c.Visited {
			return color.GreenString(s)
		}
		return s
	}
	switch c.C {
	case 'S':
		return color.RedString("S")
	case '-':
		return greenIfVisited("=")
	case '|':
		return greenIfVisited("║")
	case 'F':
		return greenIfVisited("╔")
	case '7':
		return greenIfVisited("╗")
	case 'L':
		return greenIfVisited("╚")
	case 'J':
		return greenIfVisited("╝")
	}
	return string(c.C)
}

type Connectors int

const (
	Left Connectors = 1 << iota
	Right
	Top
	Bottom
)

// Connectors returns the directions in which the cell is able to connect to other cells.
func (c Char) Connectors() Connectors {
	switch c.C {
	case 'S':
		return Top | Bottom | Left | Right
	case '-':
		return Left | Right
	case '|':
		return Top | Bottom
	case 'F':
		return Right | Bottom
	case '7':
		return Left | Bottom
	case 'L':
		return Top | Right
	case 'J':
		return Top | Left
	}
	return 0
}

func (c Char) ConnectsTo(conn Connectors) bool {
	return c.Connectors()&conn != 0
}

/*
Adjacent cells:
x 1 x
2 x 3
x 4 x
*/

func printMatrix(m [][]Char) {
	for _, line := range m {
		for _, char := range line {
			fmt.Print(char)
		}
		fmt.Println()
	}
}

func constructMatrix(lines []string) (matrix [][]Char, startX, startY int) {
	matrix = make([][]Char, len(lines))
	for y, line := range lines {
		for x, char := range line {
			if char == 'S' {
				startX, startY = x, y
			}
			matrix[y] = append(matrix[y], Char{C: char})
		}
	}
	return matrix, startX, startY
}

func solution1(ctx day.Context, lines []string) (any, error) {
	m, startX, startY := constructMatrix(lines)
	// printMatrix(m)

	steps := 0
	var previousMove Connectors
	x, y := startX, startY
	for x != startX || y != startY || steps == 0 {
		// Get the adjacent cells (Top, Left, Right, Bottom)
		adj := util.ConnectingCells(m, x, y, new(Char))
		// Test all four directions whether they connect to the current cell
		current := m[y][x]
		m[y][x].Visited = true
		// Top
		// 1. Check if current cell connects to top
		// 2. Check if adjacent cell to the top connects to bottom
		//    and-notting the previous move (as to don't go back to where we came from,
		//    i.e. if the previous move was to move down, don't move up)
		if current.ConnectsTo(Top) && adj[0].ConnectsTo(Bottom&^previousMove) {
			previousMove = Top
			y--
		} else
		// Left
		if current.ConnectsTo(Left) && adj[1].ConnectsTo(Right&^previousMove) {
			previousMove = Left
			x--
		} else
		// Right
		if current.ConnectsTo(Right) && adj[2].ConnectsTo(Left&^previousMove) {
			previousMove = Right
			x++
		} else
		// Bottom
		if current.ConnectsTo(Bottom) && adj[3].ConnectsTo(Top&^previousMove) {
			previousMove = Bottom
			y++
		}
		steps++
	}

	// printMatrix(m)

	return int(steps / 2), nil
}

func solution2(ctx day.Context, lines []string) (any, error) {
	m, startX, startY := constructMatrix(lines)
	// printMatrix(m)

	steps := 0
	var previousMove Connectors
	x, y := startX, startY
	for x != startX || y != startY || steps == 0 {
		// Get the adjacent cells
		adj := util.ConnectingCells(m, x, y, new(Char))
		// Test all four directions whether they connect to the current cell
		current := m[y][x]
		m[y][x].Visited = true
		// Top
		if current.ConnectsTo(Top) && adj[0].ConnectsTo(Bottom&^previousMove) {
			previousMove = Top
			y--
		} else
		// Left
		if current.ConnectsTo(Left) && adj[1].ConnectsTo(Right&^previousMove) {
			previousMove = Left
			x--
		} else
		// Right
		if current.ConnectsTo(Right) && adj[2].ConnectsTo(Left&^previousMove) {
			previousMove = Right
			x++
		} else
		// Bottom
		if current.ConnectsTo(Bottom) && adj[3].ConnectsTo(Top&^previousMove) {
			previousMove = Bottom
			y++
		}
		steps++
	}

	contained := 0
	for _, line := range m {
		crossed := 0
		for x, char := range line {
			if char.Visited && char.ConnectsTo(Top) {
				crossed++
			}
			if !char.Visited {
				if crossed%2 == 1 {
					char.Contained = true
					line[x] = char
					contained++
				}
			}
		}
	}
	// printMatrix(m)
	return contained, nil
}

func solution22(ctx day.Context, lines []string) (any, error) {
	m, startX, startY := constructMatrix(lines)
	// printMatrix(m)

	steps := 0
	var previousMove Connectors
	x, y := startX, startY
	var path [][2]int
	for x != startX || y != startY || steps == 0 {
		path = append(path, [2]int{x, y})
		// Get the adjacent cells
		adj := util.ConnectingCells(m, x, y, new(Char))
		// Test all four directions whether they connect to the current cell
		current := m[y][x]
		m[y][x].Visited = true
		// Top
		if current.ConnectsTo(Top) && adj[0].ConnectsTo(Bottom&^previousMove) {
			previousMove = Top
			y--
		} else
		// Left
		if current.ConnectsTo(Left) && adj[1].ConnectsTo(Right&^previousMove) {
			previousMove = Left
			x--
		} else
		// Right
		if current.ConnectsTo(Right) && adj[2].ConnectsTo(Left&^previousMove) {
			previousMove = Right
			x++
		} else
		// Bottom
		if current.ConnectsTo(Bottom) && adj[3].ConnectsTo(Top&^previousMove) {
			previousMove = Bottom
			y++
		}
		steps++
	}

	// Using shoelace formula to calculate the area of the path
	// And Pick's Theorem to calculate the number of contained points
	area := util.ShoelaceArea(path)

	return area - float64(len(path))/2 + 1, nil
}
