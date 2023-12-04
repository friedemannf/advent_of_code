package day2

import (
	"strconv"
	"strings"

	"github.com/friedemannf/advent_of_code/day"
)

const (
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
)

func init() {
	day.Register(2023, 2, day.Day{
		Solution1: []day.Solution{solution1},
		Solution2: []day.Solution{solution2},
	})
}

func solution1(ctx day.Context, lines []string) (any, error) {
	sum := 0
	for _, line := range lines {
		line = strings.TrimPrefix(line, "Game ")
		split := strings.Split(line, ": ")
		game, err := strconv.ParseInt(split[0], 10, 64)
		if err != nil {
			panic(err)
		}
		line = split[1]
		draws := strings.Split(line, "; ")
		possible := true
	impossible:
		for _, draw := range draws {
			split := strings.Split(draw, ", ")
			for _, color := range split {
				parts := strings.Split(color, " ")
				number, err := strconv.ParseInt(parts[0], 10, 64)
				if err != nil {
					panic(err)
				}
				switch parts[1] {
				case "red":
					if number > maxRed {
						possible = false
						break impossible
					}
				case "green":
					if number > maxGreen {
						possible = false
						break impossible
					}
				case "blue":
					if number > maxBlue {
						possible = false
						break impossible
					}
				}

			}
		}
		if possible {
			sum += int(game)
		}
	}
	return sum, nil
}

func solution2(ctx day.Context, lines []string) (any, error) {
	sum := 0
	for _, line := range lines {
		line = strings.TrimPrefix(line, "Game ")
		split := strings.Split(line, ": ")
		line = split[1]
		draws := strings.Split(line, "; ")
		var minRed, minGreen, minBlue int64
		for _, draw := range draws {
			split := strings.Split(draw, ", ")
			for _, color := range split {
				parts := strings.Split(color, " ")
				number, err := strconv.ParseInt(parts[0], 10, 64)
				if err != nil {
					panic(err)
				}
				switch parts[1] {
				case "red":
					if number > minRed {
						minRed = number
					}
				case "green":
					if number > minGreen {
						minGreen = number
					}
				case "blue":
					if number > minBlue {
						minBlue = number
					}
				}

			}
		}
		sum += int(minRed) * int(minGreen) * int(minBlue)
	}
	return sum, nil
}
