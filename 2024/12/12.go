package day12

import (
	"github.com/friedemannf/advent_of_code/day"
	"github.com/friedemannf/advent_of_code/util"
)

func init() {
	day.Register(2024, 12, day.Day{
		Solution1: []day.Solution{solution1},
		Solution2: []day.Solution{solution2},
	})
}

func solution1(ctx day.Context, lines []string) (any, error) {
	matrix := util.MatrixFromLines(lines)
	if ctx.Debug {
		matrix.ColorPrint(nil)
	}

	covered := util.MakeMatrix[bool](matrix.Height(), matrix.Width())
	sum := 0
	for coordinate, v := range matrix.Iterate() {
		if ok, _ := covered.Get(coordinate); ok {
			continue
		}
		perimeter := 0
		cells := matrix.FloodFill(coordinate)
		for _, cell := range cells {
			covered.Set(cell, true)
			connectingCells := matrix.ConnectingCells(cell, util.New(' '))
			for _, connectingCell := range connectingCells {
				if connectingCell != v {
					perimeter++
				}
			}
		}
		sum += len(cells) * perimeter
	}
	return sum, nil
}

func solution2(ctx day.Context, lines []string) (any, error) {
	matrix := util.MatrixFromLines(lines)
	if ctx.Debug {
		matrix.ColorPrint(nil)
	}

	covered := util.MakeMatrix[bool](matrix.Height(), matrix.Width())
	sum := 0
	for coordinate, v := range matrix.Iterate() {
		if ok, _ := covered.Get(coordinate); ok {
			continue
		}
		sides := 0
		cells := matrix.FloodFill(coordinate)
		for _, cell := range cells {
			covered.Set(cell, true)
			// 0 1 2
			// 3 X 4
			// 5 6 7
			connectingCells := matrix.AdjacentCells(cell, util.New(' '))
			// Just like for solution1, probe every direction to check if we're at a border
			// But in addition to that, check if the direction we're probing is a continuation
			// of an already existing border. For probing X upwards, there could be an:
			//   inside corner  outside corner   no corner
			//     [3] != A       [0] == A
			// [0]→  B B B          A┃B B          B B B
			// [3]→  B┏━━━          A┗━━━          ━━━━━
			//       B┃A↑A          A X↑A          A X↑A
			//       A┃A A          A A A          A A A
			//
			// Up
			if connectingCells[1] != v {
				if connectingCells[3] != v || connectingCells[0] == v {
					sides++
				}
			}
			// Left
			if connectingCells[3] != v {
				if connectingCells[6] != v || connectingCells[5] == v {
					sides++
				}
			}
			// Right
			if connectingCells[4] != v {
				if connectingCells[1] != v || connectingCells[2] == v {
					sides++
				}
			}
			// Down
			if connectingCells[6] != v {
				if connectingCells[4] != v || connectingCells[7] == v {
					sides++
				}
			}
		}
		sum += len(cells) * sides
	}
	return sum, nil
}
