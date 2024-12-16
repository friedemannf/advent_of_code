package day11

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/friedemannf/advent_of_code/day"
	"github.com/friedemannf/advent_of_code/util"
)

func init() {
	day.Register(2024, 11, day.Day{
		Solution1: []day.Solution{solution1},
		Solution2: []day.Solution{solution2},
	})
}

func solution1(ctx day.Context, lines []string) (any, error) {
	iterations := 25
	if ctx.Input != "" {
		iterations, _ = strconv.Atoi(ctx.Input)
	}
	split := strings.Split(lines[0], " ")
	stones := make([]int, len(split))
	for i, s := range split {
		stone, _ := strconv.Atoi(s)
		stones[i] = stone
	}
	for i := range iterations {
		if ctx.Debug {
			fmt.Println(i)
			fmt.Println(stones)
		}
		stones = blink(stones)
	}
	if ctx.Debug {
		fmt.Println(stones)
	}
	return len(stones), nil
}

func blink(stones []int) []int {
	out := make([]int, 0, len(stones)*2)
	for _, stone := range stones {
		switch {
		case stone == 0:
			out = append(out, 1)
		case util.Digits(stone)%2 == 0:
			power := util.Power(10, util.Digits(stone)/2)
			left := stone / power
			right := stone % power
			out = append(out, left, right)
		default:
			out = append(out, stone*2024)
		}
	}
	return out
}

func solution2(ctx day.Context, lines []string) (any, error) {
	iterations := 75
	if ctx.Input != "" {
		iterations, _ = strconv.Atoi(ctx.Input)
	}
	split := strings.Split(lines[0], " ")
	stones := make(map[int64]int64)
	for _, s := range split {
		stone, _ := strconv.ParseInt(s, 10, 64)
		stones[stone] = 1
	}
	for i := range iterations {
		if ctx.Debug {
			fmt.Println(i)
			fmt.Println(stones)
		}
		newStones := make(map[int64]int64, len(stones)*2)
		for stone, count := range stones {
			if stone < 0 {
				panic("overflow")
			}
			switch {
			case stone == 0:
				newStones[1] += count
			case util.Digits(stone)%2 == 0:
				power := util.Power(int64(10), util.Digits(stone)/2)
				left := stone / power
				right := stone % power
				newStones[left] += count
				newStones[right] += count
			default:
				newStones[stone*2024] += count
			}
		}
		stones = newStones
	}
	if ctx.Debug {
		fmt.Println(stones)
	}
	var sum int64
	for _, v := range stones {
		sum += v
	}
	return sum, nil
}
