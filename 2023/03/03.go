package day3

import (
	"github.com/friedemannf/advent_of_code/day"
	"github.com/friedemannf/advent_of_code/util"
)

func init() {
	day.Register(2023, 3, day.Day{
		Solution1: []day.Solution{solution1},
		Solution2: []day.Solution{solution2},
	})
}

func solution1(ctx day.Context, lines []string) (any, error) {
	sum := 0
	for y, line := range lines {
		currPart := 0
		adjacent := false
		for x, char := range line {
			// If digit is a number 0-9
			if char >= '0' && char <= '9' {
				// If part is already marked as adjacent, no need to check again
				if !adjacent {
					// Get adjacent cells and check for symbol
					adjacentCells := adjacentCells(lines, x, y)
					for _, adjacentCell := range adjacentCells {
						// If adjacent cell is not a number or a dot, mark as adjacent
						if !(adjacentCell >= '0' && adjacentCell <= '9') && adjacentCell != '.' {
							adjacent = true
							break
						}
					}
				}
				// Update running part with last digit
				digit := int(char - '0')
				currPart = currPart*10 + digit
			}
			// Digit is not a number or at end of line - check if this has finished a part number and update sum if part was adjacent.
			if !(char >= '0' && char <= '9') || x == len(line)-1 {
				if currPart > 0 {
					if adjacent {
						sum += currPart
						// ctx.Print(color.GreenString("%v", currPart))
					} else {
						// ctx.Print(color.WhiteString("%v", currPart))
					}
					currPart = 0
					adjacent = false
				}
				if char == '.' {
					// ctx.Print(color.HiBlackString("%c", char))
				} else if x != len(line)-1 {
					// ctx.Print(color.RedString("%c", char))
				}
			}
		}
		// ctx.Print("\n")
	}
	return sum, nil
}

func adjacentCells(lines []string, x, y int) []uint8 {
	res := make([]uint8, 0, 8)
	// 5 3 6
	// 1 X 2
	// 7 4 8

	// 1
	if x > 0 {
		res = append(res, lines[y][x-1])
	}
	// 2
	if x < len(lines[y])-1 {
		res = append(res, lines[y][x+1])
	}
	// 3
	if y > 0 {
		res = append(res, lines[y-1][x])
	}
	// 4
	if y < len(lines)-1 {
		res = append(res, lines[y+1][x])
	}
	// 5
	if x > 0 && y > 0 {
		res = append(res, lines[y-1][x-1])
	}
	// 6
	if x < len(lines[y])-1 && y > 0 {
		res = append(res, lines[y-1][x+1])
	}
	// 7
	if x > 0 && y < len(lines)-1 {
		res = append(res, lines[y+1][x-1])
	}
	// 8
	if x < len(lines[y])-1 && y < len(lines)-1 {
		res = append(res, lines[y+1][x+1])
	}
	return res
}

type partnumber int

func solution2(ctx day.Context, lines []string) (any, error) {
	// Doing two passes here
	// First pass: create a grid of pointers to part numbers
	// Second pass: iterate over grid and for every gear, check for adjacent parts
	var grid util.Matrix[*partnumber]
	// Holds the current part number/nil
	// When reaching the end of a part number (either because of a non-digit or end of line), this is reset to nil
	var currPart *partnumber
	for y, line := range lines {
		grid = append(grid, make([]*partnumber, len(line)))
		for x, char := range line {
			// If digit is a number 0-9
			if char >= '0' && char <= '9' {
				if currPart == nil {
					zero := partnumber(0)
					currPart = &zero
				}
				digit := int(char - '0')
				*currPart = (*currPart)*10 + partnumber(digit)
				grid[y][x] = currPart
			}
			// Digit is not a number, reset current part
			if !(char >= '0' && char <= '9') {
				currPart = nil
			}
		}
	}
	/*ctx.Print("Pointers:\n")
	  for _, line := range grid {
	  	for _, part := range line {
	  		if part != nil {
	  			ctx.Print(color.GreenString("X"))
	  		} else {
	  			ctx.Print(color.WhiteString("."))
	  		}
	  	}
	  	ctx.Print("\n")
	  }*/
	sum := 0
	for y, line := range lines {
		for x, char := range line {
			if char == '*' {
				adjacentCells := grid.AdjacentCells(util.Coordinate{X: x, Y: y}, nil)
				adjacentParts := make(map[*partnumber]*partnumber)
				for _, adjacentCell := range adjacentCells {
					if adjacentCell != nil {
						adjacentParts[adjacentCell] = adjacentCell
					}
				}
				// if there are two adjacent parts, multiply their numbers and add to sum
				if len(adjacentParts) == 2 {
					product := 1
					for _, part := range adjacentParts {
						product *= int(*part)
					}
					sum += product
				}
			}
		}
	}
	return sum, nil
}
