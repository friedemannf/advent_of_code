package day1

import (
	"github.com/friedemannf/advent_of_code/day"
)

func init() {
	day.Register(2015, 01, day.Day{
		Solution1: []day.Solution{solution1},
		Solution2: []day.Solution{solution2},
	})
}

func solution1(_ day.Context, lines []string) (any, error) {
	floor := 0
	for _, instruction := range lines[0] {
		switch instruction {
		case '(':
			floor++
		case ')':
			floor--
		}
	}
	return floor, nil
}

func solution2(_ day.Context, lines []string) (any, error) {
	floor := 0
	for i, instruction := range lines[0] {
		switch instruction {
		case '(':
			floor++
		case ')':
			floor--
		}
		if floor == -1 {
			return i + 1, nil
		}
	}
	return floor, nil
}
