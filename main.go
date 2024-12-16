package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/fatih/color"

	"github.com/friedemannf/advent_of_code/day"
	"github.com/friedemannf/advent_of_code/util"

	_ "github.com/friedemannf/advent_of_code/2015"
	_ "github.com/friedemannf/advent_of_code/2023"
	_ "github.com/friedemannf/advent_of_code/2024"
)

var (
	file      = flag.String("file", "input.txt", "input file")
	d         = flag.Int("day", -1, "day to run")
	y         = flag.Int("year", time.Now().Year(), "year to run")
	part      = flag.Int("part", 1, "part to run")
	iteration = flag.Int("i", 0, "iteration to run")
	benchmark = flag.Int("bench", 1, "benchmark iterations")
	debug     = flag.Bool("debug", false, "debug output")
	input     = flag.String("input", "", "optional input passed to the implementation")
)

func main() {
	red := color.New(color.FgRed).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	_ = red
	flag.Parse()

	color.Red("Advent of Code")
	if *d == -1 {
		*d = time.Now().Day()
	}

	// Read the input file.
	fmt.Printf("Reading input file %s\n", yellow(*file))
	lines, err := util.ReadLines(*file)
	if err != nil {
		panic(err)
	}

	fmt.Println("Year", yellow(*y))
	fmt.Println("Day", yellow(*d))
	d, ok := day.GetDay(*y, *d)
	if !ok {
		panic("day not found")
	}

	fmt.Println("Part", yellow(*part))
	fmt.Println("Iteration", yellow(*iteration))

	var s day.Solution
	switch *part {
	case 1:
		s = d.Solution1[*iteration]
	case 2:
		s = d.Solution2[*iteration]
	}

	start := time.Now()
	var res any
	for i := 0; i < *benchmark; i++ {
		res, err = s(day.Context{
			Context: context.Background(),
			Debug:   *debug,
			Input:   *input,
		}, lines)
	}
	dur := time.Now().Sub(start)
	if err != nil {
		panic(err)
	}

	dur = time.Duration(dur.Nanoseconds() / int64(*benchmark))

	fmt.Printf("Result: %v\n", green(res))
	color.HiBlack("Took: %s\n", dur)
}
