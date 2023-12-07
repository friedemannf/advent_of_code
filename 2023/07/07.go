package day7

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/friedemannf/advent_of_code/day"
)

func init() {
	day.Register(2023, 7, day.Day{
		Solution1: []day.Solution{solution1},
		Solution2: []day.Solution{solution2},
	})
}

type Card int

const (
	CJack Card = iota
	C1
	C2
	C3
	C4
	C5
	C6
	C7
	C8
	C9
	CT
	CJ
	CQ
	CK
	CA
)

type typ int

const (
	Five typ = iota
	Four
	FullHouse
	Three
	Two
)

type Hand struct {
	Cards [5]Card
	// Five, Four, Full House, Three, Two
	Types [5]int
	Bid   int
}

func (h Hand) String() string {
	b := strings.Builder{}
	b.WriteString("[")
	for i, c := range h.Cards {
		switch c {
		case CJack:
			b.WriteString("J")
		case C1:
			b.WriteString("1")
		case C2:
			b.WriteString("2")
		case C3:
			b.WriteString("3")
		case C4:
			b.WriteString("4")
		case C5:
			b.WriteString("5")
		case C6:
			b.WriteString("6")
		case C7:
			b.WriteString("7")
		case C8:
			b.WriteString("8")
		case C9:
			b.WriteString("9")
		case CT:
			b.WriteString("T")
		case CJ:
			b.WriteString("J")
		case CQ:
			b.WriteString("Q")
		case CK:
			b.WriteString("K")
		case CA:
			b.WriteString("A")
		}
		if i < len(h.Cards)-1 {
			b.WriteString(" ")
		}
	}
	b.WriteString("]  ")
	b.WriteString(fmt.Sprintf("%v  %d", h.Types, h.Bid))
	return b.String()
}

// Returns -1 if h < h2, 0 if h == h2, 1 if h > h2.
func CmpCards(h1, h2 Hand) int {
	for i, t1 := range h1.Types {
		t2 := h2.Types[i]
		if t1 == t2 && t2 == 0 {
			continue
		}
		if t1 < t2 {
			return -1
		} else if t1 > t2 {
			return 1
		}
		break
	}
	for i, c1 := range h1.Cards {
		c2 := h2.Cards[i]
		if c1 < c2 {
			return -1
		} else if c1 > c2 {
			return 1
		}
	}
	return 0
}

func solution1(ctx day.Context, lines []string) (any, error) {
	hands := make([]Hand, len(lines))
	for i, line := range lines {
		s := strings.Fields(line)
		m := make(map[int32]int)
		for ii, c := range s[0] {
			m[c]++
			switch c {
			case 'A':
				hands[i].Cards[ii] = CA
			case 'K':
				hands[i].Cards[ii] = CK
			case 'Q':
				hands[i].Cards[ii] = CQ
			case 'J':
				hands[i].Cards[ii] = CJ
			case 'T':
				hands[i].Cards[ii] = CT
			default:
				hands[i].Cards[ii] = Card(c - '0')
			}
		}
		hands[i].Bid, _ = strconv.Atoi(s[1])
		for _, v := range m {
			switch v {
			case 5:
				hands[i].Types[Five]++
			case 4:
				hands[i].Types[Four]++
			case 3:
				hands[i].Types[Three]++
			case 2:
				hands[i].Types[Two]++
			}
		}
		if hands[i].Types[Three] > 0 && hands[i].Types[Two] > 0 {
			hands[i].Types[FullHouse]++
		}
	}
	slices.SortFunc(hands, CmpCards)
	sum := 0
	for i, v := range hands {
		sum += v.Bid * (i + 1)
	}

	return sum, nil
}

func solution2(ctx day.Context, lines []string) (any, error) {
	hands := make([]Hand, len(lines))
	for i, line := range lines {
		s := strings.Fields(line)
		m := make(map[Card]int)
		for ii, c := range s[0] {
			switch c {
			case 'A':
				hands[i].Cards[ii] = CA
			case 'K':
				hands[i].Cards[ii] = CK
			case 'Q':
				hands[i].Cards[ii] = CQ
			case 'J':
				// Joker instead of Jack
				hands[i].Cards[ii] = CJack
			case 'T':
				hands[i].Cards[ii] = CT
			default:
				hands[i].Cards[ii] = Card(c - '0')
			}
			m[hands[i].Cards[ii]]++
		}
		hands[i].Bid, _ = strconv.Atoi(s[1])
		for c, v := range m {
			if c == CJack {
				continue
			}
			switch v {
			case 5:
				hands[i].Types[Five]++
			case 4:
				hands[i].Types[Four]++
			case 3:
				hands[i].Types[Three]++
			case 2:
				hands[i].Types[Two]++
			}
		}
		if hands[i].Types[Three] > 0 && hands[i].Types[Two] > 0 {
			hands[i].Types[FullHouse]++
			hands[i].Types[Three]--
			hands[i].Types[Two]--
		}

		// Apply jokers
		jokers := m[CJack]
		for jokers > 0 {
			jokers--
			// Fives
			if hands[i].Types[Five] > 0 {
				continue
			}
			// Fours
			if hands[i].Types[Four] > 0 {
				hands[i].Types[Five]++
				hands[i].Types[Four]--
				continue
			}
			// Full House
			if hands[i].Types[FullHouse] > 0 {
				hands[i].Types[Four]++
				hands[i].Types[FullHouse]--
				continue
			}
			// Threes
			if hands[i].Types[Three] > 0 {
				hands[i].Types[Four]++
				hands[i].Types[Three]--
				continue
			}
			// Twos
			if hands[i].Types[Two] == 1 {
				hands[i].Types[Three]++
				hands[i].Types[Two]--
				continue
			} else if hands[i].Types[Two] > 1 {
				hands[i].Types[FullHouse]++
				hands[i].Types[Two] -= 2
				continue
			}
			hands[i].Types[Two]++
		}

	}
	slices.SortFunc(hands, CmpCards)
	sum := 0
	for i, v := range hands {
		sum += v.Bid * (i + 1)
		// fmt.Println(v)
	}

	return sum, nil
}
