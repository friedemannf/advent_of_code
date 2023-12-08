package day8

import (
	"regexp"

	"github.com/friedemannf/advent_of_code/day"
)

func init() {
	day.Register(2023, 8, day.Day{
		Solution1: []day.Solution{solution1},
		Solution2: []day.Solution{solution2},
	})
}

type Node struct {
	V    string
	L, R *Node
}

func solution1(ctx day.Context, lines []string) (any, error) {
	r := regexp.MustCompile(`^(...) = \((...), (...)\)$`)
	nodes := make(map[string]*Node)
	instructions := ""
	for i, line := range lines {
		if i == 0 {
			instructions = line
			continue
		}
		if matches := r.FindStringSubmatch(line); len(matches) == 4 {
			node, ok := nodes[matches[1]]
			if !ok {
				node = &Node{V: matches[1]}
				nodes[matches[1]] = node
			}
			left, ok := nodes[matches[2]]
			if !ok {
				left = &Node{V: matches[2]}
				nodes[matches[2]] = left
			}
			right, ok := nodes[matches[3]]
			if !ok {
				right = &Node{V: matches[3]}
				nodes[matches[3]] = right
			}
			node.L = left
			node.R = right
		}
	}
	node := nodes["AAA"]
	target := nodes["ZZZ"]
	i := 0
	for node != target {
		inst := instructions[i%len(instructions)]
		// fmt.Println(node.V, string(inst))
		switch inst {
		case 'L':
			node = node.L
		case 'R':
			node = node.R
		}
		i++
	}
	return i, nil
}

func solution2(ctx day.Context, lines []string) (any, error) {
	r := regexp.MustCompile(`^(...) = \((...), (...)\)$`)
	nodes := make(map[string]*Node)
	var (
		startingNodes []*Node
	)
	endingNodes := make(map[*Node]struct{})
	instructions := ""
	for i, line := range lines {
		if i == 0 {
			instructions = line
			continue
		}
		if matches := r.FindStringSubmatch(line); len(matches) == 4 {
			node, ok := nodes[matches[1]]
			if !ok {
				node = &Node{V: matches[1]}
				nodes[matches[1]] = node
			}
			left, ok := nodes[matches[2]]
			if !ok {
				left = &Node{V: matches[2]}
				nodes[matches[2]] = left
			}
			right, ok := nodes[matches[3]]
			if !ok {
				right = &Node{V: matches[3]}
				nodes[matches[3]] = right
			}
			node.L = left
			node.R = right
			if node.V[2] == 'A' {
				startingNodes = append(startingNodes, node)
			}
			if node.V[2] == 'Z' {
				endingNodes[node] = struct{}{}
			}
		}
	}
	// fmt.Println(startingNodes)
	// fmt.Println(endingNodes)
	var lengths []int
	for _, node := range startingNodes {
		l := loopLength(node, endingNodes, instructions)
		lengths = append(lengths, l)
	}
	// fmt.Println(lengths)
	lcm := LCM(lengths...)
	return lcm, nil
}

func loopLength(node *Node, targetNodes map[*Node]struct{}, instructions string) int {
	i := 0
	for {
		if _, ok := targetNodes[node]; ok {
			return i
		}
		// fmt.Println(node.V, node.L.V, node.R.V, i)
		inst := instructions[i%len(instructions)]
		// fmt.Println(node.V, string(inst))
		switch inst {
		case 'L':
			node = node.L
		case 'R':
			node = node.R
		}
		i++
	}
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(ints ...int) int {
	result := ints[0] * ints[1] / GCD(ints[0], ints[1])

	for i := 0; i < len(ints)-2; i++ {
		result = LCM(result, ints[i+2])
	}

	return result
}
