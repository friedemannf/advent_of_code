package day13

import (
	"math/bits"

	"github.com/friedemannf/advent_of_code/day"
)

func init() {
	day.Register(2023, 13, day.Day{
		Solution1: []day.Solution{solution1, solution11},
		Solution2: []day.Solution{solution2},
	})
}

func solution1(_ day.Context, lines []string) (any, error) {
	lines = append(lines, "")
	sum := 0
	var (
		rows, columns []string
	)
	for _, line := range lines {
		if len(line) == 0 {
			// Checking horizontally
			if match, ok := checkMatch(rows); ok {
				sum += 100 * match
			}
			// Checking vertically
			if match, ok := checkMatch(columns); ok {
				sum += match
			}

			rows = rows[:0]
			columns = columns[:0]
			continue
		}

		rows = append(rows, line)
		if len(columns) < len(line) {
			columns = make([]string, len(line))
		}
		for i, c := range line {
			columns[i] += string(c)
		}
	}
	return sum, nil
}

func checkMatch[T comparable](strings []T) (int, bool) {
	for i := 1; i < len(strings); i++ {
		match := true
		for j := 1; j < len(strings); j++ {
			ix1 := i - j
			ix2 := i + j - 1

			if ix1 < 0 {
				break
			}
			if ix2 >= len(strings) {
				break
			}

			s1 := strings[ix1]
			s2 := strings[ix2]

			if s1 != s2 {
				match = false
				break
			}

		}
		if match {
			return i, true
		}
	}
	return 0, false
}

func solution11(_ day.Context, lines []string) (any, error) {
	lines = append(lines, "")
	sum := 0
	var (
		rows, columns []uint64
	)
	for _, line := range lines {
		if len(line) == 0 {
			// Checking horizontally
			if match, ok := checkMatch(rows); ok {
				sum += 100 * match
			}
			// Checking vertically
			if match, ok := checkMatch(columns); ok {
				sum += match
			}

			rows = rows[:0]
			columns = columns[:0]
			continue
		}

		if len(columns) < len(line) {
			columns = make([]uint64, len(line))
		}
		row := uint64(0)
		for i, c := range line {
			columns[i] = columns[i] << 1
			if c == '#' {
				row += 1 << i
				columns[i] = columns[i] + 1
			}
		}
		rows = append(rows, row)
	}
	return sum, nil
}

func checkMatches(rows []uint64) []int {
	var matches []int
	for i := 1; i < len(rows); i++ {
		diff := 0
		for j := 1; j < len(rows); j++ {
			ix1 := i - j
			ix2 := i + j - 1

			if ix1 < 0 {
				break
			}
			if ix2 >= len(rows) {
				break
			}

			r1 := rows[ix1]
			r2 := rows[ix2]

			diff += bits.OnesCount64(r1 ^ r2)
			if diff > 1 {
				break
			}
		}
		if diff <= 1 {
			matches = append(matches, i)
		}
	}
	return matches
}

func solution2(_ day.Context, lines []string) (any, error) {
	lines = append(lines, "")
	sum := 0
	var (
		rows, columns []uint64
	)
	for _, line := range lines {
		if len(line) == 0 {
			// Checking horizontally
			matchBefore, ok := checkMatch(rows)
			matchesAfter := checkMatches(rows)

			var diff []int
			for _, m := range matchesAfter {
				if ok && m == matchBefore {
					continue
				}
				diff = append(diff, m)
			}
			if len(diff) == 1 {
				sum += 100 * diff[0]
			}

			// Checking vertically
			matchBefore, ok = checkMatch(columns)
			matchesAfter = checkMatches(columns)

			diff = diff[:0]
			for _, m := range matchesAfter {
				if ok && m == matchBefore {
					continue
				}
				diff = append(diff, m)
			}
			if len(diff) == 1 {
				sum += diff[0]
			}

			rows = rows[:0]
			columns = columns[:0]
			continue
		}

		if len(columns) < len(line) {
			columns = make([]uint64, len(line))
		}
		row := uint64(0)
		for i, c := range line {
			columns[i] = columns[i] << 1
			if c == '#' {
				row += 1 << i
				columns[i] = columns[i] + 1
			}
		}
		rows = append(rows, row)
	}
	return sum, nil
}
