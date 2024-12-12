package day2

import (
	"strconv"
	"strings"

	"github.com/friedemannf/advent_of_code/day"
	"github.com/friedemannf/advent_of_code/util"
)

func init() {
	day.Register(2015, 02, day.Day{
		Solution1: []day.Solution{solution1},
		Solution2: []day.Solution{solution2},
	})
}

func solution1(_ day.Context, lines []string) (any, error) {
	area := 0
	for _, present := range lines {
		s := strings.Split(present, "x")
		l, w, h := util.Must(strconv.Atoi(s[0])), util.Must(strconv.Atoi(s[1])), util.Must(strconv.Atoi(s[2]))
		area += 2*l*w + 2*w*h + 2*h*l
		if l >= w {
			if l >= h {
				area += w * h
			} else {
				area += l * w
			}
		} else {
			if w >= h {
				area += l * h
			} else {
				area += l * w
			}
		}
	}
	return area, nil
}

func solution2(_ day.Context, lines []string) (any, error) {
	ribbon := 0
	for _, present := range lines {
		s := strings.Split(present, "x")
		l, w, h := util.Must(strconv.Atoi(s[0])), util.Must(strconv.Atoi(s[1])), util.Must(strconv.Atoi(s[2]))
		ribbon += l * w * h
		if l >= w {
			if l >= h {
				ribbon += 2*w + 2*h
			} else {
				ribbon += 2*l + 2*w
			}
		} else {
			if w >= h {
				ribbon += 2*l + 2*h
			} else {
				ribbon += 2*l + 2*w
			}
		}
	}
	return ribbon, nil
}
