package day8

import (
	"github.com/fatih/color"

	"github.com/friedemannf/advent_of_code/day"
	"github.com/friedemannf/advent_of_code/util"
)

func init() {
	day.Register(2024, 8, day.Day{
		Solution1: []day.Solution{solution1},
		Solution2: []day.Solution{solution2},
	})
}

func solution1(_ day.Context, lines []string) (any, error) {
	antennas := make(map[rune][]util.Coordinate)
	matrix := util.Matrix[rune](make([][]rune, len(lines)))
	for posY, line := range lines {
		matrix[posY] = []rune(line)
		for posX, char := range line {
			if char == '.' {
				continue
			}
			antennas[char] = append(antennas[char], util.Coordinate{
				X: posX,
				Y: posY,
			})
		}
	}
	antinodes := make(map[util.Coordinate]bool)
	for _, coordinates := range antennas {
		combinations := util.Combinations(coordinates, 2)
		for _, combination := range combinations {
			from, to := combination[0], combination[1]
			vec := from.DistanceTo(to)

			antinode1 := to.Add(vec)
			antinode2 := from.Sub(vec)
			if matrix.Set(antinode1, '#') {
				antinodes[antinode1] = true
			}
			if matrix.Set(antinode2, '#') {
				antinodes[antinode2] = true
			}
		}
	}
	matrix.ColorPrint(map[rune]color.Attribute{
		'#': color.FgGreen,
		'.': color.FgWhite,
		0:   color.FgBlue, // default color
	})
	return len(antinodes), nil
}

func solution2(_ day.Context, lines []string) (any, error) {
	antennas := make(map[rune][]util.Coordinate)
	matrix := util.Matrix[rune](make([][]rune, len(lines)))
	for posY, line := range lines {
		matrix[posY] = []rune(line)
		for posX, char := range line {
			if char == '.' {
				continue
			}
			antennas[char] = append(antennas[char], util.Coordinate{
				X: posX,
				Y: posY,
			})
		}
	}
	antinodes := make(map[util.Coordinate]bool)
	for _, coordinates := range antennas {
		combinations := util.Combinations(coordinates, 2)
		for _, combination := range combinations {
			from, to := combination[0], combination[1]
			vec := from.DistanceTo(to)

			// Other than in part 1, the from/to antennas are antinodes themselves
			// By adding to the from node and subtracting from the to node, we get the to/from nodes
			tmp := from.Copy()
			for {
				// Repeat until the point isn't on the matrix anymore
				tmp = tmp.Add(vec)
				if matrix.Set(tmp, '#') {
					antinodes[tmp] = true
				} else {
					break
				}
			}
			tmp = to.Copy()
			for {
				tmp = tmp.Sub(vec)
				if matrix.Set(tmp, '#') {
					antinodes[tmp] = true
				} else {
					break
				}
			}
		}
	}
	matrix.ColorPrint(map[rune]color.Attribute{
		'#': color.FgGreen,
		'.': color.FgWhite,
		0:   color.FgBlue, // default color
	})
	return len(antinodes), nil
}
