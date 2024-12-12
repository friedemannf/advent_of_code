package day9

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/friedemannf/advent_of_code/color"
	"github.com/friedemannf/advent_of_code/day"
)

func init() {
	day.Register(2024, 9, day.Day{
		Solution1: []day.Solution{solution1},
		Solution2: []day.Solution{solution2},
	})
}

func solution1(_ day.Context, lines []string) (any, error) {
	line := lines[0]
	var diskmap []int
	fmt.Println(line)
	for i, char := range line {
		n := char - '0'
		if i%2 == 0 {
			// file
			for range n {
				diskmap = append(diskmap, i/2)
			}
		} else {
			// empty space
			for range n {
				diskmap = append(diskmap, -1)
			}
		}
	}
	start := 0
	end := len(diskmap) - 1
	for {
		if start == end {
			break
		}
		if diskmap[start] != -1 {
			start++
			continue
		}
		if diskmap[end] == -1 {
			end--
			continue
		}
		diskmap[start], diskmap[end] = diskmap[end], diskmap[start]
	}
	sum := 0
	for i, char := range diskmap {
		if char == -1 {
			continue
		}
		sum += i * char
	}
	return sum, nil
}

type block struct {
	Empty  bool
	Length int32
	ID     int
}

func (b block) String() string {
	s := strings.Builder{}
	for range b.Length {
		if b.Empty {
			s.WriteString(".")
		} else {
			s.WriteString(strconv.Itoa(b.ID))
		}
	}
	return s.String()
}

func solution2(ctx day.Context, lines []string) (any, error) {
	line := lines[0]

	var diskmap []block
	for i, char := range line {
		n := char - '0'
		diskmap = append(diskmap, block{
			Empty:  i%2 != 0,
			Length: n,
			ID:     i / 2,
		})
	}
	debugPrint := func(start, end int) {
		if !ctx.Debug {
			return
		}
		for i, b := range diskmap {
			if i == start {
				fmt.Print(color.Green(b))
			} else if i == end {
				fmt.Print(color.Red(b))
			} else {
				fmt.Print(b)
			}
		}
		fmt.Println()

	}
	highestID := diskmap[len(diskmap)-1].ID + 1
	elementToMove := len(diskmap) - 1
	for elementToMove >= 0 {
		debugPrint(0, elementToMove)
		if diskmap[elementToMove].Empty || diskmap[elementToMove].ID >= highestID {
			elementToMove--
			debugPrint(0, elementToMove)
			continue
		}
		// find place
		start := 0
		for start < elementToMove && !diskmap[start].Empty || diskmap[start].Length < diskmap[elementToMove].Length {
			start++
			debugPrint(start, elementToMove)
			continue
		}
		if start < elementToMove {
			debugPrint(start, elementToMove)
			// switch
			highestID = diskmap[elementToMove].ID
			diskmap[start].Length -= diskmap[elementToMove].Length
			diskmap = slices.Insert(diskmap, start, diskmap[elementToMove])
			diskmap[elementToMove+1].Empty = true
			debugPrint(start, elementToMove)
		}
		elementToMove--
	}
	sum := 0
	iter := 0
	for _, char := range diskmap {
		if char.Empty {
			iter += int(char.Length)
			continue
		}
		for n := range char.Length {
			sum += (int(n) + iter) * char.ID
		}
		iter += int(char.Length)
	}
	return sum, nil
}
