package day18

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/friedemannf/advent_of_code/day"
	"github.com/friedemannf/advent_of_code/util"
)

func init() {
	day.Register(2023, 18, day.Day{
		Solution1: []day.Solution{solution1, solution11},
		Solution2: []day.Solution{solution2},
	})
}

func solution1(_ day.Context, lines []string) (any, error) {
	m := util.Matrix[rune]{}
	currCoord := util.Coordinate{
		X: 0,
		Y: 0,
	}
	m.Set(currCoord, 'X')
	for _, line := range lines {
		split := strings.Split(line, " ")
		distance, _ := strconv.Atoi(split[1])
		switch split[0] {
		case "R":
			for x := 0; x < distance; x++ {
				currCoord.X++
				m.Set(currCoord, '#')
			}
		case "L":
			for x := 0; x < distance; x++ {
				currCoord.X--
				m.Set(currCoord, '#')
			}
		case "U":
			for x := 0; x < distance; x++ {
				currCoord.Y--
				m.Set(currCoord, '#')
			}
		case "D":
			for x := 0; x < distance; x++ {
				currCoord.Y++
				m.Set(currCoord, '#')
			}
		}
	}
	// fmt.Println(m)
	// fmt.Println()

	sum := 0
	previousY := 0
	previousV := 0
	incoming := 0
	crossed := 0
	for coord, v := range m.Iterate() {
		if coord.Y != previousY {
			// New line
			previousY = coord.Y
			previousV = 0
			incoming = 0
			crossed = 0
		}
		if v == '#' {
			sum++
			cells := m.ConnectingCells(coord, nil)
			if cells[0] == '#' {
				incoming--
			}
			if cells[3] == '#' {
				incoming++
			}
		}
		if v == 0 {
			if previousV == '#' {
				if incoming == 0 {
					crossed++
				}
				incoming = 0
			}
			if crossed%2 == 1 {
				m.Set(coord, 'X')
				sum++
			}
		}
		previousV = int(v)
	}
	// fmt.Println(m)
	return sum, nil
}

func solution11(_ day.Context, lines []string) (any, error) {
	var path [][2]int
	currX, currY := 0, 0
	pathLen := 0
	for _, line := range lines {
		split := strings.Split(line, " ")
		distance, _ := strconv.Atoi(split[1])
		pathLen += distance
		switch split[0] {
		case "R":
			currX += distance
			path = append(path, [2]int{currX, currY})
		case "L":
			currX -= distance
			path = append(path, [2]int{currX, currY})
		case "U":
			currY -= distance
			path = append(path, [2]int{currX, currY})
		case "D":
			currY += distance
			path = append(path, [2]int{currX, currY})
		}
	}
	area := util.ShoelaceArea(path)
	// Pick's theorem to calculate the number of contained points
	// = area - pathLen/2 + 1
	// adding the path length to the area to account for the points on the path
	// = area + pathLen/2 + 1
	return area + float64(pathLen)/2 + 1, nil
}

var r = regexp.MustCompile(`^. \d+ \(#(.....)(.)\)$`)

func solution2(_ day.Context, lines []string) (any, error) {
	var path [][2]int64
	var (
		currX, currY int64
		pathLen      int64
	)
	for _, line := range lines {
		matches := r.FindStringSubmatch(line)
		distance, _ := strconv.ParseInt(matches[1], 16, 64)
		pathLen += distance
		switch matches[2] {
		case "0":
			currX += distance
			path = append(path, [2]int64{currX, currY})
		case "2":
			currX -= distance
			path = append(path, [2]int64{currX, currY})
		case "3":
			currY -= distance
			path = append(path, [2]int64{currX, currY})
		case "1":
			currY += distance
			path = append(path, [2]int64{currX, currY})
		}
	}
	area := util.ShoelaceArea(path)
	// Pick's theorem to calculate the number of contained points
	// = area - pathLen/2 + 1
	// adding the path length to the area to account for the points on the path
	// = area + pathLen/2 + 1
	return strconv.FormatInt(int64(area+float64(pathLen)/2+1), 10), nil
}
