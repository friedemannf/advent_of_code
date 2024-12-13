package day10

import (
	"github.com/friedemannf/advent_of_code/day"
	"github.com/friedemannf/advent_of_code/util"
)

func init() {
	day.Register(2024, 10, day.Day{
		Solution1: []day.Solution{solution1},
		Solution2: []day.Solution{solution2},
	})
}

func solution1(_ day.Context, lines []string) (any, error) {
	matrix := util.MatrixFromLines(lines)
	trailheads := make(map[util.Coordinate]map[util.Coordinate]bool)
	for k, v := range matrix.Iterate() {
		if v == '0' {
			trailheads[k] = make(map[util.Coordinate]bool)
		}
	}
	sum := 0
	for trailhead, summits := range trailheads {
		for _, summit := range walk(matrix, trailhead) {
			// There are multiple paths leading to the same summit, de-duplicating them
			summits[summit] = true
		}
		sum += len(summits)
	}
	return sum, nil
}

func walk(matrix util.Matrix[rune], coordinate util.Coordinate) []util.Coordinate {
	v, _ := matrix.Get(coordinate)
	if v-'0' == 9 {
		return []util.Coordinate{coordinate}
	}
	var res []util.Coordinate
	next := matrix.ConnectingCells(coordinate, util.New(' '))
	// Up
	if next[0]-v == 1 { // if next is 1 step up from the current coordinate
		res = append(res, walk(matrix, coordinate.Add(util.VecUp))...)
	}
	// Left
	if next[1]-v == 1 {
		res = append(res, walk(matrix, coordinate.Add(util.VecLeft))...)
	}
	// Right
	if next[2]-v == 1 {
		res = append(res, walk(matrix, coordinate.Add(util.VecRight))...)
	}
	// Down
	if next[3]-v == 1 {
		res = append(res, walk(matrix, coordinate.Add(util.VecDown))...)
	}
	return res
}

func solution2(ctx day.Context, lines []string) (any, error) {
	matrix := util.MatrixFromLines(lines)
	trailheads := make(map[util.Coordinate]map[util.Coordinate]bool)
	for k, v := range matrix.Iterate() {
		if v == '0' {
			trailheads[k] = make(map[util.Coordinate]bool)
		}
	}
	sum := 0
	for trailhead := range trailheads {
		sum += len(walk(matrix, trailhead))
	}
	return sum, nil
}
