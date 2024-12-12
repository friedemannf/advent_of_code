package day7

import (
	"fmt"
	"sync"

	"github.com/brunoga/deep"

	"github.com/friedemannf/advent_of_code/color"
	"github.com/friedemannf/advent_of_code/day"
)

func init() {
	day.Register(2024, 7, day.Day{
		Solution1: []day.Solution{solution1},
		Solution2: []day.Solution{solution2},
	})
}

func solution1(_ day.Context, lines []string) (any, error) {
	mapp := make([][]rune, len(lines))
	posY, posX, dir := 0, 0, 0
	for i, line := range lines {
		mapp[i] = make([]rune, len(line))
		for ii, c := range line {
			mapp[i][ii] = c
			if c == '^' {
				posY, posX = i, ii
			}
		}
	}
	sum := 0
	for posY >= 0 && posY < len(mapp) && posX >= 0 && posX < len(mapp[0]) {
		was := &mapp[posY][posX]
		if *was == '#' {
			// Hit obstacle -- go back and turn right
			switch dir {
			case 0:
				posY++
			case 1:
				posX--
			case 2:
				posY--
			case 3:
				posX++
			}
			dir = (dir + 1) % 4
			continue
		}
		if *was != 'X' {
			sum++
			*was = 'X'
		}
		switch dir {
		case 0:
			// North
			posY--
		case 1:
			// East
			posX++
		case 2:
			// South
			posY++
		case 3:
			// West
			posX--
		}
	}
	for _, runes := range mapp {
		fmt.Println(string(runes))
	}
	return sum, nil
}

func solution2(_ day.Context, lines []string) (any, error) {
	mapp := make([][]rune, len(lines))
	posY, posX := 0, 0
	for i, line := range lines {
		mapp[i] = make([]rune, len(line))
		for ii, c := range line {
			mapp[i][ii] = c
			if c == '^' {
				posY, posX = i, ii
			}
		}
	}
	sum := 0
	wg := sync.WaitGroup{}
	mux := sync.Mutex{}
	for y, line := range mapp {
		for x, r := range line {
			switch r {
			case '^', '>', 'v', '<', '#':
			default:
				newMapp := deep.MustCopy(mapp)
				wg.Add(1)
				go func() {
					defer wg.Done()
					newMapp[y][x] = 'O'
					if checkMap(newMapp, posY, posX) {
						// printMap(newMapp)
						mux.Lock()
						sum++
						mux.Unlock()
					}

				}()
				// return sum, nil
			}
		}
	}
	wg.Wait()
	return sum, nil
}

func checkMap(mapp [][]rune, posY, posX int) bool {
	dir := 0
	hit := false
	visited := make(map[struct{ x, y, dir int }]bool)
	for {
		//printMapPosition(mapp, posY, posX, dir)
		// Check if the guard left the map
		if posY < 0 || posY >= len(mapp) || posX < 0 || posX >= len(mapp[0]) {
			return false
		}
		if visited[struct{ x, y, dir int }{x: posX, y: posY, dir: dir}] {
			return true
		}
		visited[struct{ x, y, dir int }{x: posX, y: posY, dir: dir}] = true
		was := &mapp[posY][posX]
		switch *was {
		// Check if we're in a loop
		case '|':
			if dir%2 == 1 {
				// Travelling orthogonally - crossing previous path
				*was = '+'
			}
		case '-':
			if dir%2 == 0 {
				*was = '+'
			}
		case '+':
			if hit {
				// Had to backtrack in the previous iteration, no loop
				hit = false
			}
		case '#', 'O':
			// Hit obstacle -- go back and turn right
			switch dir {
			case 0:
				posY++
			case 1:
				posX--
			case 2:
				posY--
			case 3:
				posX++
			}
			mapp[posY][posX] = '+'
			dir = (dir + 1) % 4
			hit = true
			continue
		default:
			// Just mark the field as visited
			if dir%2 == 0 {
				// North/South
				*was = '|'
			} else {
				// East/West
				*was = '-'
			}
		}
		// Move to next field
		switch dir {
		case 0:
			// North
			posY--
		case 1:
			// East
			posX++
		case 2:
			// South
			posY++
		case 3:
			// West
			posX--
		}
	}

}

func printMapPosition(mapp [][]rune, posY, posX, dir int) {
	if posY >= 0 && posY < len(mapp) && posX >= 0 && posX < len(mapp[0]) {
		was := mapp[posY][posX]
		guard := ' '
		switch dir {
		case 0:
			guard = '^'
		case 1:
			guard = '>'
		case 2:
			guard = 'v'
		case 3:
			guard = '<'
		}
		mapp[posY][posX] = guard
		printMap(mapp)
		mapp[posY][posX] = was
	} else {
		printMap(mapp)
	}
}

func printMap(mapp [][]rune) {
	for _, line := range mapp {
		for _, r := range line {
			switch r {
			case '|', '-', '+', 'X':
				fmt.Print(color.Red(string(r)))
			case '#':
				fmt.Print(color.Green(string(r)))
			case '^', '>', 'v', '<':
				fmt.Print(color.Blue(string(r)))
			case 'O':
				fmt.Print(color.Yellow(string(r)))
			default:
				fmt.Print(string(r))
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
