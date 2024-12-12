package day1

import (
	"slices"
	"strconv"
	"strings"

	"github.com/friedemannf/advent_of_code/day"
	"github.com/friedemannf/advent_of_code/util"
)

func init() {
	day.Register(2024, 1, day.Day{
		Solution1: []day.Solution{solution1},
		Solution2: []day.Solution{solution2},
	})
}

func solution1(_ day.Context, lines []string) (any, error) {
	left := make([]int, len(lines))
	right := make([]int, len(lines))
	var err error

	for i, line := range lines {
		split := strings.Split(line, "   ")
		left[i], err = strconv.Atoi(split[0])
		if err != nil {
			return nil, err
		}
		right[i], err = strconv.Atoi(split[1])
		if err != nil {
			return nil, err
		}

	}
	slices.Sort(left)
	slices.Sort(right)

	var sum int
	for i, left := range left {
		right := right[i]
		sum += util.Abs(left - right)
	}
	return sum, nil
}

func solution2(_ day.Context, lines []string) (any, error) {
	left := make([]int, len(lines))
	right := make(map[int]int, len(lines))
	var err error

	for i, line := range lines {
		split := strings.Split(line, "   ")
		left[i], err = strconv.Atoi(split[0])
		if err != nil {
			return nil, err
		}
		r, err := strconv.Atoi(split[1])
		if err != nil {
			return nil, err
		}
		right[r]++
	}

	var sum int
	for _, n := range left {
		sum += right[n] * n
	}
	return sum, nil
}
