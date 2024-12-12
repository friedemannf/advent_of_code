package day2

import (
	"strconv"
	"strings"
	"sync"

	"github.com/friedemannf/advent_of_code/day"
	"github.com/friedemannf/advent_of_code/util"
)

func init() {
	day.Register(2024, 2, day.Day{
		Solution1: []day.Solution{solution1},
		Solution2: []day.Solution{solution2, solution22},
	})
}

func checkReport(levels []int) bool {
	var prev *int
	var direction int
	safe := true
	for _, n := range levels {
		if prev == nil {
			prev = &n // Requires go 1.22
			continue
		}
		switch direction {
		case 0:
			// unknown direction
			if *prev < n {
				direction = 1
			} else if *prev > n {
				direction = -1
			} else {
				safe = false
				break
			}
		case 1:
			if *prev >= n {
				safe = false
				break
			}
		case -1:
			if *prev <= n {
				safe = false
				break
			}
		}

		if util.Abs(*prev-n) > 3 {
			safe = false
			break
		}
		prev = &n
	}
	return safe
}

func solution1(_ day.Context, lines []string) (any, error) {
	sum := 0
	for _, report := range lines {
		s := strings.Split(report, " ")
		levels := make([]int, len(s))
		for i, level := range s {
			n, _ := strconv.Atoi(level)
			levels[i] = n
		}
		if checkReport(levels) {
			sum++
		}
	}
	return sum, nil
}

func solution2(_ day.Context, lines []string) (any, error) {
	sum := 0
	for _, report := range lines {
		s := strings.Split(report, " ")
		levels := make([]int, len(s))
		for i, level := range s {
			n, _ := strconv.Atoi(level)
			levels[i] = n
		}
		if checkReport(levels) {
			sum++
			continue
		}
		temp := make([]int, len(levels)-1)
		for i := 0; i < len(levels); i++ {
			copy(temp, levels[:i])
			temp = temp[:i]
			temp = append(temp, levels[i+1:]...)
			// fmt.Println(i, " ", levels, "->", levels[:i], " + ", levels[i+1:], " : ", temp)
			if checkReport(temp) {
				sum++
				break
			}
		}
	}
	return sum, nil
}

func solution22(_ day.Context, lines []string) (any, error) {
	sum := 0
	res := make(chan any)
	wg := sync.WaitGroup{}
	for _, report := range lines {
		wg.Add(1)
		go func(report string) {
			defer wg.Done()
			s := strings.Split(report, " ")
			levels := make([]int, len(s))
			for i, level := range s {
				n, _ := strconv.Atoi(level)
				levels[i] = n
			}
			if checkReport(levels) {
				res <- true
				return
			}
			temp := make([]int, len(levels)-1)
			for i := 0; i < len(levels); i++ {
				copy(temp, levels[:i])
				temp = temp[:i]
				temp = append(temp, levels[i+1:]...)
				// fmt.Println(i, " ", levels, "->", levels[:i], " + ", levels[i+1:], " : ", temp)
				if checkReport(temp) {
					res <- true
					break
				}
			}
		}(report)
	}
	go func() {
		wg.Wait()
		close(res)
	}()
	for range res {
		sum++
	}
	return sum, nil
}
