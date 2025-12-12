package main

import (
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
			if conn == "" {
				continue
			}
			g[parts[0]] = append(g[parts[0]], strings.TrimSpace(conn))
		}
	}
	return &g
}

func memo[M comparable, T any](f func(func(M, T) int) func(M, T) int) func(M, T) int {
	memo := map[M]int{}
	var self func(M, T) int

	self = func(m M, t T) int {
		if val, ok := memo[m]; ok {
			return val
		}

		memo[m] = f(self)(m, t)
		return memo[m]
	}

	return self
}

func Out(self func(string, *graph) int) func(string, *graph) int {
	return func(entry string, g *graph) int {
		if entry == "out" {
			return 1
		}

		sum := 0
		for _, e := range (*g)[entry] {
			sum += self(e, g)
		}

		return sum
	}
}

type args struct {
	entry string
	fft   bool
	dac   bool
}

func OutPlusVisit(self func(args, *graph) int) func(args, *graph) int {
	return func(a args, g *graph) int {
		if a.entry == "out" {
			if a.fft && a.dac {
				return 1
			}
			return 0
		}

		sum := 0
		for _, e := range (*g)[a.entry] {
				sum += self(args{e, a.fft || a.entry == "fft", a.dac || a.entry == "dac"}, g)
		}

		return sum
	}
}

func part1(file string) int {
	g := NewGraph(file)
	o := memo(Out)
	return o("you", g)
}

func part2(file string) int {
	g := NewGraph(file)
	o := memo(OutPlusVisit)
	return o(args{"svr", false, false}, g)
}

func main() {
	fmt.Printf("part1(\"./inputs/day-11\"): %v\n", part1("./inputs/day-11"))
	fmt.Printf("part2(\"./inputs/day-11\"): %v\n", part2("./inputs/day-11"))
}
