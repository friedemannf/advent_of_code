package util

func AdjacentCells[T any](lines [][]T, x, y int, oob *T) []T {
	res := make([]T, 0, 8)
	// 1 2 3
	// 4 X 5
	// 6 7 8

	appendIfOutOfBounds := func() {
		if oob != nil {
			res = append(res, *oob)
		}
	}

	// 1
	if x > 0 && y > 0 {
		res = append(res, lines[y-1][x-1])
	} else {
		appendIfOutOfBounds()
	}
	// 2
	if y > 0 {
		res = append(res, lines[y-1][x])
	} else {
		appendIfOutOfBounds()
	}
	// 3
	if x < len(lines[y])-1 && y > 0 {
		res = append(res, lines[y-1][x+1])
	} else {
		appendIfOutOfBounds()
	}

	// 4
	if x > 0 {
		res = append(res, lines[y][x-1])
	} else {
		appendIfOutOfBounds()
	}
	// 5
	if x < len(lines[y])-1 {
		res = append(res, lines[y][x+1])
	} else {
		appendIfOutOfBounds()
	}

	// 6
	if x > 0 && y < len(lines)-1 {
		res = append(res, lines[y+1][x-1])
	} else {
		appendIfOutOfBounds()
	}
	// 7
	if y < len(lines)-1 {
		res = append(res, lines[y+1][x])
	} else {
		appendIfOutOfBounds()
	}
	// 8
	if x < len(lines[y])-1 && y < len(lines)-1 {
		res = append(res, lines[y+1][x+1])
	} else {
		appendIfOutOfBounds()
	}
	return res
}

func ConnectingCells[T any](lines [][]T, x, y int, oob *T) []T {
	res := make([]T, 0, 4)
	// x 1 x
	// 2 X 3
	// x 4 x

	appendIfOutOfBounds := func() {
		if oob != nil {
			res = append(res, *oob)
		}
	}

	// 1
	if y > 0 {
		res = append(res, lines[y-1][x])
	} else {
		appendIfOutOfBounds()
	}
	// 2
	if x > 0 {
		res = append(res, lines[y][x-1])
	} else {
		appendIfOutOfBounds()
	}
	// 3
	if x < len(lines[y])-1 {
		res = append(res, lines[y][x+1])
	} else {
		appendIfOutOfBounds()
	}
	// 4
	if y < len(lines)-1 {
		res = append(res, lines[y+1][x])
	} else {
		appendIfOutOfBounds()
	}
	return res
}
