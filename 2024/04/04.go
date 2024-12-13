package day4

import (
	"fmt"
	"strings"

	"github.com/friedemannf/advent_of_code/day"
	"github.com/friedemannf/advent_of_code/util"
)

func init() {
	day.Register(2024, 4, day.Day{
		Solution1: []day.Solution{solution1},
		Solution2: []day.Solution{solution2},
	})
}

func solution1(_ day.Context, lines []string) (any, error) {
	matrix := util.Matrix[rune](make([][]rune, len(lines)))
	for i, line := range lines {
		matrix[i] = []rune(line)
	}
	dirs := []util.Modifier{
		util.Horizontal, util.Horizontal | util.Reverse,
		util.Vertical, util.Vertical | util.Reverse,
		util.DiagonalLR, util.DiagonalLR | util.Reverse,
		util.DiagonalRL, util.DiagonalRL | util.Reverse,
	}
	sum := 0
	for _, dir := range dirs {
		t := matrix.Transform(dir)
		for _, runes := range t {
			n := strings.Count(string(runes), "XMAS")
			fmt.Println(string(runes), n)
			sum += n
		}
		fmt.Println("---")
	}
	return sum, nil
}

func solution2(_ day.Context, lines []string) (any, error) {
	matrix := util.Matrix[rune](make([][]rune, len(lines)))
	for i, line := range lines {
		matrix[i] = []rune(line)
	}
	sum := 0
	for y, line := range matrix {
		for x, c := range line {
			if c == 'A' {
				ac := matrix.AdjacentCells(util.Coordinate{X: x, Y: y}, util.New(' '))
				if (string([]rune{ac[0], ac[7]}) == "MS" || string([]rune{ac[7], ac[0]}) == "MS") &&
					(string([]rune{ac[5], ac[2]}) == "MS" || string([]rune{ac[2], ac[5]}) == "MS") {
					sum++
				}
			}
		}
	}
	return sum, nil
}
