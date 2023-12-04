package day4

import (
	"strings"

	"github.com/friedemannf/advent_of_code/day"
	"github.com/friedemannf/advent_of_code/util"
)

func init() {
	day.Register(2023, 4, day.Day{
		Solution1: []day.Solution{solution1},
		Solution2: []day.Solution{solution2},
	})
}

func solution1(ctx day.Context, lines []string) (any, error) {
	sum := 0
	for _, line := range lines {
		line = strings.Split(line, ": ")[1]
		parts := strings.Split(line, " | ")
		winning := strings.Fields(parts[0])
		got := strings.Fields(parts[1])
		match := 0
		for _, v := range winning {
			for _, g := range got {
				if v == g {
					match++
					break
				}
			}
		}
		if match > 0 {
			sum += util.Power(2, match-1)
		}
	}
	return sum, nil
}

func solution2(ctx day.Context, lines []string) (any, error) {
	sum := 0
	cards := make([]int, len(lines))
	for y, line := range lines {
		cards[y]++
		line = strings.Split(line, ": ")[1]
		parts := strings.Split(line, " | ")
		winning := strings.Fields(parts[0])
		got := strings.Fields(parts[1])
		match := 0
		for _, v := range winning {
			for _, g := range got {
				if v == g {
					match++
					break
				}
			}
		}
		for i := 1; i <= match && len(cards) > (y+i); i++ {
			cards[y+i] += cards[y]
		}
	}
	for _, v := range cards {
		sum += v
	}
	return sum, nil
}
