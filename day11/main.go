package main

import (
	// "fmt"
	"fmt"
	"strings"

	"github.com/markome-beep/AOC-2025/shared"
)

type graph = map[string][]string

func NewGraph(file string) *graph {
	g := make(map[string][]string)
	for line := range shared.ReadLines(file, "\n") {
		parts := strings.Split(line, ":")
		g[parts[0]] = make([]string, 0)
		for conn := range strings.SplitSeq(parts[1], " ") {
			g[parts[0]] = append(g[parts[0]], strings.TrimSpace(conn))
		}
	}
	return &g
}

func Out(entry string, memo *map[string]int, g *graph) int {
	if entry == "out" {
		return 1
	}
	if val, ok := (*memo)[entry]; ok {
		return val
	}

	sum := 0
	for _, e := range (*g)[entry] {
		sum += Out(e, memo, g)
	}

	return sum
}

func OutPlusVisit(entry string, memo *map[string]int, visited *map[string]bool, g *graph, memoOut *map[string]int) int {
	if entry == "out" {
		return 0
	}

	if val, ok := (*memo)[entry]; ok {
		return val
	}

	(*visited)[entry] = true
	defer delete(*visited, entry)

	sum := 0
	for _, e := range (*g)[entry] {
		if (*visited)["fft"] && (*visited)["dac"] {
			sum += Out(e, memoOut, g)
		} else {
			sum += OutPlusVisit(e, memo, visited, g, memoOut)
		}
	}

	return sum
}

func part1(file string) int {
	g := NewGraph(file)
	memo := make(map[string]int)
	return Out("you", &memo, g)
}

func part2(file string) int {
	g := NewGraph(file)
	memo := make(map[string]int)
	memoOut := make(map[string]int)

	return OutPlusVisit("svr", &memo, &map[string]bool{}, g, &memoOut)
}

func main() {
	fmt.Printf("part1(\"./inputs/day-11\"): %v\n", part1("./inputs/day-11"))
	fmt.Printf("part2(\"./inputs/day-11\"): %v\n", part2("./inputs/day-11"))
}
