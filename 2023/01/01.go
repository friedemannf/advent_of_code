package day1

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/friedemannf/advent_of_code/day"
)

func init() {
	day.Register(2023, 1, day.Day{
		Solution1: []day.Solution{solution1, solution11},
		Solution2: []day.Solution{solution2, solution22},
	})
}

func solution1(ctx day.Context, lines []string) (any, error) {
	sum := 0
	for _, line := range lines {
		line = strings.Trim(line, "abcdefghijklmnopqrstuvwxyz")
		concatted := fmt.Sprintf("%s%s", string(line[0]), string(line[len(line)-1]))
		number, err := strconv.ParseInt(concatted, 10, 64)
		if err != nil {
			return nil, err
		}
		sum += int(number)
	}
	return sum, nil
}

func solution11(ctx day.Context, lines []string) (any, error) {
	sum := 0
	for _, line := range lines {
		var rune1, rune2 rune
		// From the beginning of the string, find the first number.
		for i := 0; i < len(line); i++ {
			if line[i] >= '0' && line[i] <= '9' {
				rune1 = rune(line[i])
				break
			}
		}
		// From the end of the string, find the first number.
		for i := len(line) - 1; i >= 0; i-- {
			if line[i] >= '0' && line[i] <= '9' {
				rune2 = rune(line[i])
				break
			}
		}
		number := fmt.Sprintf("%s%s", string(rune1), string(rune2))
		integer, err := strconv.ParseInt(number, 10, 64)
		if err != nil {
			return nil, err
		}
		sum += int(integer)
	}
	return sum, nil
}

func solution2(ctx day.Context, lines []string) (any, error) {
	sum := 0
	for _, line := range lines {
		var rune1, rune2 rune
		// From the beginning of the string, find the first number.
		for i := 0; i < len(line); i++ {
			if line[i] >= '0' && line[i] <= '9' {
				rune1 = rune(line[i])
				break
			}
			if n := findNumber(line[i:]); n != 0 {
				rune1 = n
				break
			}
		}
		// From the end of the string, find the first number.
		for i := len(line) - 1; i >= 0; i-- {
			if line[i] >= '0' && line[i] <= '9' {
				rune2 = rune(line[i])
				break
			}
			if n := findNumber(line[i:]); n != 0 {
				rune2 = n
				break
			}
		}
		number := fmt.Sprintf("%s%s", string(rune1), string(rune2))
		integer, err := strconv.ParseInt(number, 10, 64)
		if err != nil {
			return nil, err
		}
		sum += int(integer)
	}
	return sum, nil
}

func solution22(ctx day.Context, lines []string) (any, error) {
	ch := make(chan int64)
	wg := sync.WaitGroup{}
	for _, line := range lines {
		wg.Add(1)
		go func(line string) {

			var rune1, rune2 rune
			// From the beginning of the string, find the first number.
			for i := 0; i < len(line); i++ {
				if line[i] >= '0' && line[i] <= '9' {
					rune1 = rune(line[i])
					break
				}
				if n := findNumber(line[i:]); n != 0 {
					rune1 = n
					break
				}
			}
			// From the end of the string, find the first number.
			for i := len(line) - 1; i >= 0; i-- {
				if line[i] >= '0' && line[i] <= '9' {
					rune2 = rune(line[i])
					break
				}
				if n := findNumber(line[i:]); n != 0 {
					rune2 = n
					break
				}
			}
			number := fmt.Sprintf("%s%s", string(rune1), string(rune2))
			integer, err := strconv.ParseInt(number, 10, 64)
			if err != nil {
				panic(err)
			}
			ch <- integer
			wg.Done()
		}(line)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	sum := int64(0)
	for n := range ch {
		sum += n
	}
	return sum, nil
}

func findNumber(s string) rune {
	if len(s) <= 3 {
		return 0
	}
	switch s[0] {
	case 'o', 't', 'f', 's', 'e', 'n':
	default:
		return 0
	}
	numbers := map[string]rune{
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}
	for k, v := range numbers {
		if strings.HasPrefix(s, k) {
			return v
		}
	}
	return 0
}
