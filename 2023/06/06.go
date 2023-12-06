package day6

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/friedemannf/advent_of_code/day"
)

func init() {
	day.Register(2023, 6, day.Day{
		Solution1: []day.Solution{solution1(calcRecords1), solution1(calcRecords2)},
		Solution2: []day.Solution{solution2(calcRecords1), solution2(calcRecords2)},
	})
}

func calcRecords1(t, record int) int {
	num := 0
	for i := 0; i < t; i++ {
		if calcDistance(i, t) > record {
			num++
		}
	}
	return num
}

func calcRecords2(t, record int) int {
	num1 := 0
	num2 := 0
	for i := 0; i < t; i++ {
		if calcDistance(i, t) > record {
			num1 = i
			break
		}
	}
	for i := t; i >= 0; i-- {
		if calcDistance(i, t) > record {
			num2 = i
			break
		}
	}
	return num2 - num1 + 1
}

func solution1(calc func(int, int) int) day.Solution {
	return func(ctx day.Context, lines []string) (any, error) {
		times := strings.Fields(lines[0])
		records := strings.Fields(lines[1])

		sum := 1
		for i := 1; i < len(times); i++ {
			time, _ := strconv.Atoi(times[i])
			record, _ := strconv.Atoi(records[i])
			sum *= calc(time, record)
		}

		return sum, nil
	}
}

func calcDistance(hold, time int) int {
	return (time - hold) * hold
}

func solution2(calc func(int, int) int) day.Solution {
	return func(ctx day.Context, lines []string) (any, error) {
		timeS := strings.Builder{}
		recordS := strings.Builder{}
		for _, char := range strings.Split(lines[0], ":")[1] {
			if !unicode.IsSpace(char) {
				timeS.WriteRune(char)
			}
		}
		for _, char := range strings.Split(lines[1], ":")[1] {
			if !unicode.IsSpace(char) {
				recordS.WriteRune(char)
			}
		}
		time, _ := strconv.Atoi(timeS.String())
		record, _ := strconv.Atoi(recordS.String())

		res := calc(time, record)

		return res, nil
	}
}
