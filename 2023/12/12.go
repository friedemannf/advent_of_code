package day12

import (
	"strconv"
	"strings"

	"github.com/friedemannf/advent_of_code/day"
)

func init() {
	day.Register(2023, 12, day.Day{
		Solution1: []day.Solution{solution1},
		Solution2: []day.Solution{},
	})
}

func prepareRow(s string) (string, []int) {
	split := strings.Split(s, " ")
	str := split[0]
	split2 := strings.Split(split[1], ",")
	ints := make([]int, len(split2))
	for i, v := range split2 {
		ints[i], _ = strconv.Atoi(v)
	}
	return str, ints
}

func isValid(s string, springs []int) bool {
	parts := strings.FieldsFunc(s, func(r rune) bool {
		return r == '.'
	})
	if len(parts) != len(springs) {
		return false
	}
	for i, part := range parts {
		if len(part) != springs[i] {
			return false
		}
	}
	return true
}

func generateArrangements(s string) []string {
	if len(s) == 0 {
		return []string{""}
	}
	c := s[0]
	switch c {
	case '.', '#':
		arr := generateArrangements(s[1:])
		return prefixAll(arr, string(c))
	case '?':
		arr := generateArrangements(s[1:])
		return prefixAll(arr, ".", "#")
	}
	return nil
}

func prefixAll(s []string, prefix ...string) []string {
	var res []string
	for _, v := range s {
		for _, p := range prefix {
			res = append(res, p+v)
		}
	}
	return res
}

func runRow1(s string) int {
	str, ints := prepareRow(s)
	arr := generateArrangements(str)
	sum := 0
	for _, a := range arr {
		if isValid(a, ints) {
			sum++
		} else {
		}
	}
	return sum
}

func solution1(ctx day.Context, lines []string) (any, error) {
	sum := 0
	for _, line := range lines {
		sum += runRow1(line)
	}
	return sum, nil
}
