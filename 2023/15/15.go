package day15

import (
	"strconv"
	"strings"

	"github.com/friedemannf/advent_of_code/day"
)

func init() {
	day.Register(2023, 15, day.Day{
		Solution1: []day.Solution{solution1},
		Solution2: []day.Solution{solution2},
	})
}

func Hash(s string) int {
	v := 0
	for _, c := range s {
		v += int(c)
		v *= 17
		v = v % 256
	}
	return v
}

func solution1(_ day.Context, lines []string) (any, error) {
	sum := 0
	for _, line := range lines {
		for _, s := range strings.Split(line, ",") {
			sum += Hash(s)
		}
	}
	return sum, nil
}

type Lens struct {
	FocalLength int
	Label       string
}

type Box struct {
	Lenses []Lens
}

func (b *Box) RemoveIfExists(label string) {
	i := 0
	for _, l := range b.Lenses {
		if l.Label != label {
			b.Lenses[i] = l
			i++
		}
	}
	b.Lenses = b.Lenses[:i]
}

func (b *Box) AddLens(l Lens) {
	for i, lens := range b.Lenses {
		if lens.Label == l.Label {
			b.Lenses[i] = l
			return
		}
	}

	b.Lenses = append(b.Lenses, l)
}

func solution2(_ day.Context, lines []string) (any, error) {
	boxes := [256]Box{}
	for _, line := range lines {
		for _, s := range strings.Split(line, ",") {
			// Addition
			if split := strings.Split(s, "="); len(split) == 2 {
				label := split[0]
				focalLength, _ := strconv.Atoi(split[1])
				box := boxes[Hash(label)]
				box.AddLens(Lens{
					FocalLength: focalLength,
					Label:       label,
				})
				boxes[Hash(label)] = box

			}
			// Removal
			if split := strings.Split(s, "-"); len(split) == 2 {
				label := split[0]
				box := boxes[Hash(label)]
				box.RemoveIfExists(label)
				boxes[Hash(label)] = box
			}
		}
	}
	sum := 0
	for i, box := range boxes {
		for ii, lens := range box.Lenses {
			sum += (i + 1) * (ii + 1) * lens.FocalLength
		}
	}
	return sum, nil
}
