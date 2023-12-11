package day11

import (
	"github.com/friedemannf/advent_of_code/day"
	"github.com/friedemannf/advent_of_code/util"
)

func init() {
	day.Register(2023, 11, day.Day{
		Solution1: []day.Solution{solution(1)},
		Solution2: []day.Solution{solution(1e6 - 1)},
	})
}

func solution(expansion int) day.Solution {
	return func(ctx day.Context, lines []string) (any, error) {
		// Used to construct offsets
		rows := make([]int, len(lines))
		columns := make([]int, len(lines[0]))

		var galaxies [][2]int
		for y, line := range lines {
			for x, c := range line {
				if c == '#' {
					rows[y]++
					columns[x]++
					galaxies = append(galaxies, [2]int{x, y})
				}
			}
		}
		running := 0
		for i, row := range rows {
			rows[i] = running
			if row == 0 {
				running += expansion
			}
		}
		running = 0
		for i, column := range columns {
			columns[i] = running
			if column == 0 {
				running += expansion
			}
		}

		// Add offsets
		for i := range galaxies {
			galaxies[i][0] += columns[galaxies[i][0]]
			galaxies[i][1] += rows[galaxies[i][1]]
		}

		// Permutate galaxies
		sum := 0
		for i := 0; i < len(galaxies); i++ {
			for j := i + 1; j < len(galaxies); j++ {
				galaxy1 := galaxies[i]
				galaxy2 := galaxies[j]
				// Calculate the Manhattan distance between the two galaxies
				dist := util.Abs(galaxy2[0]-galaxy1[0]) + util.Abs(galaxy2[1]-galaxy1[1])
				sum += dist
			}
		}
		return sum, nil
	}
}
