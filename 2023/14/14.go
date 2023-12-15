package day14

import (
	"github.com/friedemannf/advent_of_code/day"
)

func init() {
	day.Register(2023, 14, day.Day{
		Solution1: []day.Solution{solution1},
		Solution2: []day.Solution{},
	})
}

func solution1(_ day.Context, lines []string) (any, error) {
	var m [][]int32
	for y, line := range lines {
		m = append(m, nil)
		for x, c := range line {
			if c == 'O' {
				m[y] = append(m[y], '.')
				i := y - 1
				for ; i >= 0; i-- {
					p := m[i][x]
					if p != '.' {
						break
					}
				}
				m[i+1][x] = 'O'
			} else {
				m[y] = append(m[y], c)
			}
		}
	}
	sum := 0
	for y, line := range m {
		// fmt.Println(y, string(line))
		for _, c := range line {
			if c == 'O' {
				sum += len(line) - y
			}
		}
	}
	return sum, nil
}
