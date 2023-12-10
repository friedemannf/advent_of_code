package day9

import (
	"slices"
	"strconv"
	"strings"

	"github.com/friedemannf/advent_of_code/day"
)

func init() {
	day.Register(2023, 9, day.Day{
		Solution1: []day.Solution{solution1},
		Solution2: []day.Solution{solution2},
	})
}

func solution1(ctx day.Context, lines []string) (any, error) {
	sum := 0
	for _, line := range lines {
		var sequence []int
		for _, v := range strings.Fields(line) {
			i, _ := strconv.Atoi(v)
			sequence = append(sequence, i)
		}
		sum += extrapolate(sequence)[len(sequence)]
	}
	return sum, nil
}

func extrapolate(x []int) []int {
	// fmt.Println(x)
	allZeroes := true
	for _, v := range x {
		if v != 0 {
			allZeroes = false
			break
		}
	}
	if allZeroes {
		return append(x, 0)
	}

	y := make([]int, len(x)-1)
	for i := 0; i < len(y); i++ {
		y[i] = x[i+1] - x[i]
	}
	x = append(x, x[len(x)-1]+extrapolate(y)[len(y)])
	// fmt.Println(x)
	return x
}

func solution2(ctx day.Context, lines []string) (any, error) {
	sum := 0
	for _, line := range lines {
		var sequence []int
		for _, v := range strings.Fields(line) {
			i, _ := strconv.Atoi(v)
			sequence = append(sequence, i)
		}
		slices.Reverse(sequence)
		sum += extrapolate(sequence)[len(sequence)]
	}
	return sum, nil
}
