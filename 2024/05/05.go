package day5

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/friedemannf/advent_of_code/day"
)

func init() {
	day.Register(2024, 5, day.Day{
		Solution1: []day.Solution{solution1},
		Solution2: []day.Solution{solution2},
	})
}

func solution1(_ day.Context, lines []string) (any, error) {
	section := 0
	// X should be printed before Y
	lt := make(map[int][]int)
	// X should be printed after Y
	gt := make(map[int][]int)
	var updates [][]int
	for _, line := range lines {
		if line == "" {
			section++
			continue
		}
		if section == 0 {
			split := strings.Split(line, "|")
			before, _ := strconv.Atoi(split[0])
			after, _ := strconv.Atoi(split[1])
			lt[before] = append(lt[before], after)
			gt[after] = append(gt[after], before)
		} else {
			split := strings.Split(line, ",")
			temp := make([]int, len(split))
			for i, s := range split {
				temp[i], _ = strconv.Atoi(s)
			}
			updates = append(updates, temp)
		}
	}
	sum := 0
	for _, update := range updates {
		tmp := make([]int, len(update))
		copy(tmp, update)
		// cmp(a, b) should return a negative number when a < b, a positive number when
		// a > b and zero when a == b or a and b are incomparable in the sense of
		// a strict weak ordering.
		slices.SortStableFunc(update, func(a, b int) int {
			for _, x := range lt[a] {
				if x == b {
					return -1
				}
			}
			for _, x := range gt[a] {
				if x == b {
					return 1
				}
			}
			return 0
		})
		fmt.Println(update)
		if slices.Equal(tmp, update) {
			middleNumber := update[len(update)/2]
			sum += middleNumber
		}
	}
	return sum, nil
}

func solution2(_ day.Context, lines []string) (any, error) {
	section := 0
	// X should be printed before Y
	lt := make(map[int][]int)
	// X should be printed after Y
	gt := make(map[int][]int)
	var updates [][]int
	for _, line := range lines {
		if line == "" {
			section++
			continue
		}
		if section == 0 {
			split := strings.Split(line, "|")
			before, _ := strconv.Atoi(split[0])
			after, _ := strconv.Atoi(split[1])
			lt[before] = append(lt[before], after)
			gt[after] = append(gt[after], before)
		} else {
			split := strings.Split(line, ",")
			temp := make([]int, len(split))
			for i, s := range split {
				temp[i], _ = strconv.Atoi(s)
			}
			updates = append(updates, temp)
		}
	}
	sum := 0
	for _, update := range updates {
		tmp := make([]int, len(update))
		copy(tmp, update)
		// cmp(a, b) should return a negative number when a < b, a positive number when
		// a > b and zero when a == b or a and b are incomparable in the sense of
		// a strict weak ordering.
		slices.SortStableFunc(update, func(a, b int) int {
			for _, x := range lt[a] {
				if x == b {
					return -1
				}
			}
			for _, x := range gt[a] {
				if x == b {
					return 1
				}
			}
			return 0
		})
		fmt.Println(update)
		if !slices.Equal(tmp, update) {
			middleNumber := update[len(update)/2]
			sum += middleNumber
		}
	}
	return sum, nil
}
