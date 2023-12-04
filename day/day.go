package day

import (
	"context"
	"fmt"
)

type Context struct {
	context.Context
	Debug bool
}

func (c Context) PrintDebug(a ...any) {
	if c.Debug {
		fmt.Print(a...)
	}
}

type Solution func(ctx Context, lines []string) (any, error)

type Day struct {
	Solution1, Solution2 []Solution
}

// [year][day]
var solutions = map[int]map[int]Day{}

func Register(y, d int, day Day) {
	if solutions[y] == nil {
		solutions[y] = map[int]Day{}
	}
	solutions[y][d] = day
}

func GetDay(y, d int) (Day, bool) {
	day, ok := solutions[y][d]
	return day, ok
}
