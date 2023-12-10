package day10

import (
	"github.com/fatih/color"

	"github.com/friedemannf/advent_of_code/day"
	"github.com/friedemannf/advent_of_code/util"
)

func init() {
	day.Register(2023, 10, day.Day{
		Solution1: []day.Solution{solution1},
		Solution2: []day.Solution{solution2},
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
		return color.YellowString("S")
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

func solution1(ctx day.Context, lines []string) (any, error) {
	m := make([][]Char, len(lines))
	startX, startY := 0, 0
	for y, line := range lines {
		for x, char := range line {
			if char == 'S' {
				startX, startY = x, y
			}
			m[y] = append(m[y], Char{C: char})
		}
	}

	/*for _, line := range m {
	  	fmt.Println(line)
	  }
	  fmt.Println()*/

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
			// fmt.Println("Top")
			previousMove = Top
			y--
		} else
		// Left
		if current.ConnectsTo(Left) && adj[1].ConnectsTo(Right&^previousMove) {
			// fmt.Println("Left")
			previousMove = Left
			x--
		} else
		// Right
		if current.ConnectsTo(Right) && adj[2].ConnectsTo(Left&^previousMove) {
			// fmt.Println("Right")
			previousMove = Right
			x++
		} else
		// Bottom
		if current.ConnectsTo(Bottom) && adj[3].ConnectsTo(Top&^previousMove) {
			// fmt.Println("Bottom")
			previousMove = Bottom
			y++
		}
		steps++
	}

	/*for _, line := range m {
		fmt.Println(line)
	}*/

	return int(steps / 2), nil
}

func solution2(ctx day.Context, lines []string) (any, error) {
	m := make([][]Char, len(lines))
	startX, startY := 0, 0
	for y, line := range lines {
		for x, char := range line {
			if char == 'S' {
				startX, startY = x, y
			}
			m[y] = append(m[y], Char{C: char})
		}
	}

	/*for _, line := range m {
		for _, char := range line {
			fmt.Print(char)
		}
		fmt.Println()
	}*/

	// fmt.Println()
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
			// fmt.Println("Top")
			previousMove = Top
			y--
		} else
		// Left
		if current.ConnectsTo(Left) && adj[1].ConnectsTo(Right&^previousMove) {
			// fmt.Println("Left")
			previousMove = Left
			x--
		} else
		// Right
		if current.ConnectsTo(Right) && adj[2].ConnectsTo(Left&^previousMove) {
			// fmt.Println("Right")
			previousMove = Right
			x++
		} else
		// Bottom
		if current.ConnectsTo(Bottom) && adj[3].ConnectsTo(Top&^previousMove) {
			// fmt.Println("Bottom")
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
			// fmt.Print(char)
		}
		// fmt.Println()
	}

	return contained, nil
}
