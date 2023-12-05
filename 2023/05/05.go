package day5

import (
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/schollz/progressbar/v3"

	"github.com/friedemannf/advent_of_code/day"
)

func init() {
	day.Register(2023, 5, day.Day{
		Solution1: []day.Solution{solution1},
		Solution2: []day.Solution{solution2, solution22, solution23},
	})
}

type mapRange struct {
	From, End, Offset int
}

func solution1(ctx day.Context, lines []string) (any, error) {
	var seeds []int
	mapIx := map[string]int{
		"seed-to-soil map:":            0,
		"soil-to-fertilizer map:":      1,
		"fertilizer-to-water map:":     2,
		"water-to-light map:":          3,
		"light-to-temperature map:":    4,
		"temperature-to-humidity map:": 5,
		"humidity-to-location map:":    6,
	}
	var ranges [7][]mapRange
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if i == 0 {
			for _, s := range strings.Fields(strings.TrimPrefix(line, "seeds: ")) {
				seed, err := strconv.Atoi(s)
				if err != nil {
					return nil, err
				}
				seeds = append(seeds, int(seed))
			}
		}
		ix, ok := mapIx[line]
		if ok {
			i++
			for ; i < len(lines); i++ {
				if s := strings.Fields(lines[i]); len(s) == 3 {
					to, _ := strconv.Atoi(s[0])
					from, _ := strconv.Atoi(s[1])
					length, _ := strconv.Atoi(s[2])
					ranges[ix] = append(ranges[ix], mapRange{
						From:   from,
						End:    from + length - 1,
						Offset: to - from,
					})
				} else {
					break
				}
			}
		}
	}
	// Sort the ranges to be able to break early
	for _, r := range ranges {
		slices.SortFunc(r, func(i, j mapRange) int {
			return i.From - j.From
		})
	}
	for i := range seeds {
		// fmt.Print(seed)
		for _, m := range ranges {
			for _, soil := range m {
				if seeds[i] >= soil.From && seeds[i] <= soil.End {
					seeds[i] += soil.Offset
					break
				} else if seeds[i] < soil.From {
					break
				}
			}
			// fmt.Print(" -> ", seeds[i])

		}
		// // fmt.Println()
	}
	// Find min seed
	min := seeds[0]
	for _, seed := range seeds {
		if seed < min {
			min = seed
		}
	}
	return min, nil
}

func solution2(ctx day.Context, lines []string) (any, error) {
	var seeds [][2]int
	mapIx := map[string]int{
		"seed-to-soil map:":            0,
		"soil-to-fertilizer map:":      1,
		"fertilizer-to-water map:":     2,
		"water-to-light map:":          3,
		"light-to-temperature map:":    4,
		"temperature-to-humidity map:": 5,
		"humidity-to-location map:":    6,
	}
	var ranges [7][]mapRange
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if i == 0 {
			fields := strings.Fields(strings.TrimPrefix(line, "seeds: "))
			for i := 0; i < len(fields); i += 2 {
				seed, err := strconv.Atoi(fields[i])
				if err != nil {
					return nil, err
				}
				length, err := strconv.Atoi(fields[i+1])
				if err != nil {
					return nil, err
				}
				seeds = append(seeds, [2]int{int(seed), int(length)})
			}
		}
		ix, ok := mapIx[line]
		if ok {
			i++
			for ; i < len(lines); i++ {
				if s := strings.Fields(lines[i]); len(s) == 3 {
					to, _ := strconv.Atoi(s[0])
					from, _ := strconv.Atoi(s[1])
					length, _ := strconv.Atoi(s[2])
					ranges[ix] = append(ranges[ix], mapRange{
						From:   from,
						End:    from + length - 1,
						Offset: to - from,
					})
				} else {
					break
				}
			}
		}
	}
	// Sort the ranges to be able to break early
	for _, r := range ranges {
		slices.SortFunc(r, func(i, j mapRange) int {
			return i.From - j.From
		})
	}
	total := 0
	for i := 0; i < len(seeds); i++ {
		for ii := 0; ii < seeds[i][1]; ii++ {
			total++
		}
	}
	bar := progressbar.Default(int64(total), "Brute forcing...")
	min := math.MaxInt
	for i := 0; i < len(seeds); i++ {
		for ii := 0; ii < seeds[i][1]; ii++ {
			seed := seeds[i][0] + ii
			// fmt.Print("seed: ", seed)
		next:
			for _, m := range ranges {
				for _, soil := range m {
					if seed >= soil.From && seed <= soil.End {
						seed += soil.Offset
						continue next
					} else if seed < soil.From {
						break
					}
				}
			}
			// fmt.Print(" -> ", seed)
			if seed < min {
				min = seed
			}
			bar.Add(1)
			// // fmt.Println()
		}
	}
	return min, nil
}

func solution22(ctx day.Context, lines []string) (any, error) {
	var seeds []int
	mapIx := map[string]int{
		"seed-to-soil map:":            0,
		"soil-to-fertilizer map:":      1,
		"fertilizer-to-water map:":     2,
		"water-to-light map:":          3,
		"light-to-temperature map:":    4,
		"temperature-to-humidity map:": 5,
		"humidity-to-location map:":    6,
	}
	var ranges [7][]mapRange
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if i == 0 {
			fields := strings.Fields(strings.TrimPrefix(line, "seeds: "))
			for i := 0; i < len(fields); i += 2 {
				seed, err := strconv.Atoi(fields[i])
				if err != nil {
					return nil, err
				}
				length, err := strconv.Atoi(fields[i+1])
				if err != nil {
					return nil, err
				}
				for ii := 0; ii < length; ii++ {
					seeds = append(seeds, int(seed)+ii)
				}
			}
		}
		ix, ok := mapIx[line]
		if ok {
			i++
			for ; i < len(lines); i++ {
				if s := strings.Fields(lines[i]); len(s) == 3 {
					to, _ := strconv.Atoi(s[0])
					from, _ := strconv.Atoi(s[1])
					length, _ := strconv.Atoi(s[2])
					ranges[ix] = append(ranges[ix], mapRange{
						From:   from,
						End:    from + length - 1,
						Offset: to - from,
					})
				} else {
					break
				}
			}
		}
	}
	// Sort the ranges to be able to break early
	for _, r := range ranges {
		slices.SortFunc(r, func(i, j mapRange) int {
			return i.From - j.From
		})
	}
	bar := progressbar.Default(int64(len(seeds)), "Brute forcing...")
	min := math.MaxInt
	for seed := range seeds {
		// fmt.Print("seed: ", seed)
	next:
		for _, m := range ranges {
			for _, soil := range m {
				if seed >= soil.From && seed <= soil.End {
					seed += soil.Offset
					continue next
				} else if seed < soil.From {
					break
				}
			}
		}
		// fmt.Print(" -> ", seed)
		if seed < min {
			min = seed
		}
		bar.Add(1)
		// // fmt.Println()
	}

	return min, nil
}

func solution23(ctx day.Context, lines []string) (any, error) {
	var seeds [][2]int
	mapIx := map[string]int{
		"seed-to-soil map:":            0,
		"soil-to-fertilizer map:":      1,
		"fertilizer-to-water map:":     2,
		"water-to-light map:":          3,
		"light-to-temperature map:":    4,
		"temperature-to-humidity map:": 5,
		"humidity-to-location map:":    6,
	}
	var ranges [7][]mapRange
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if i == 0 {
			fields := strings.Fields(strings.TrimPrefix(line, "seeds: "))
			for i := 0; i < len(fields); i += 2 {
				seed, err := strconv.Atoi(fields[i])
				if err != nil {
					return nil, err
				}
				length, err := strconv.Atoi(fields[i+1])
				if err != nil {
					return nil, err
				}
				seeds = append(seeds, [2]int{int(seed), int(seed + length - 1)})
			}
		}
		ix, ok := mapIx[line]
		if ok {
			i++
			for ; i < len(lines); i++ {
				if s := strings.Fields(lines[i]); len(s) == 3 {
					to, _ := strconv.Atoi(s[0])
					from, _ := strconv.Atoi(s[1])
					length, _ := strconv.Atoi(s[2])
					ranges[ix] = append(ranges[ix], mapRange{
						From:   from,
						End:    from + length - 1,
						Offset: to - from,
					})
				} else {
					break
				}
			}
		}
	}
	// Sort the seeds
	slices.SortFunc(seeds, func(i, j [2]int) int {
		return i[0] - j[0]
	})
	// Sort the ranges to be able to break early
	for _, r := range ranges {
		slices.SortFunc(r, func(i, j mapRange) int {
			return i.From - j.From
		})
	}

	for _, r := range ranges {
		// fmt.Println("Seeds", seeds)
		// fmt.Println("Range", r)

		var seedsAfter [][2]int
	next:
		for _, s := range seeds {
			// fmt.Print(s, " ")
			for _, v := range r {
				// Seed smaller range
				// S  |---|
				// R        |---|
				// -> |xxx|
				if s[0] < v.From && s[1] < v.From {
					a := s
					// fmt.Println("-1>", a)
					seedsAfter = append(seedsAfter, a)
					continue next
				}
				// S  |--xxxx|
				// R     |----|
				// -> |xx|XXX|
				if s[0] < v.From && s[1] <= v.End {
					a := [][2]int{{s[0], v.From - 1}, {v.From + v.Offset, s[1] + v.Offset}}
					// fmt.Println("-2>", a)
					seedsAfter = append(seedsAfter, a...)
					continue next
				}
				// S  |---------|
				// R     |---|
				// -> |xx|XXX|--|  -> Next range
				if s[0] < v.From && s[1] > v.End {
					a := [][2]int{{s[0], v.From - 1}, {v.From + v.Offset, v.End + v.Offset}}
					// fmt.Println("-3>", a)
					seedsAfter = append(seedsAfter, a...)
					s[0] = v.End + 1
					continue
				}
				// S    |---|
				// R  |-------|
				// ->   |XXX|
				if s[0] >= v.From && s[1] <= v.End {
					a := [2]int{s[0] + v.Offset, s[1] + v.Offset}
					// fmt.Println("-4>", a)
					seedsAfter = append(seedsAfter, a)
					continue next
				}
				// S     |-----|
				// R  |-----|
				// ->    |XX|--|  -> Next range
				if s[0] <= v.End && s[1] > v.End {
					a := [][2]int{{s[0] + v.Offset, v.End + v.Offset}}
					// fmt.Println("-5>", a)
					seedsAfter = append(seedsAfter, a...)
					s[0] = v.End + 1
					continue
				}
			}
			// S       |---|
			// R |---|
			a := s
			// fmt.Println("-6>", a)
			seedsAfter = append(seedsAfter, a)
		}
		slices.SortFunc(seedsAfter, func(i, j [2]int) int {
			return i[0] - j[0]
		})
		// fmt.Println("After", seedsAfter)
		seeds = seedsAfter
		// fmt.Println("---")
	}

	return seeds[0][0], nil
}
