package main

import (
	"fmt"
	"github.com/bits-and-blooms/bitset"
)

type point struct {
	from [2]int // Dir edge came from
	to   [2]int // Dir edge left to
	x    uint
	y    uint
}

type path []*point

type grid struct {
	data *bitset.BitSet
	width uint
	heigth uint
}

func (p1 *point) area(p2 *point) uint {
	dx := max(p1.x, p2.x) - min(p1.x, p2.x) + 1
	dy := max(p1.y, p2.y) - min(p1.y, p2.y) + 1
	return dx * dy
}

func part1(file string) uint {
	ps := parsePoints(file)
	ps.scale()
	maxArea := uint(0)

	for i, p1 := range ps {
		for _, p2 := range ps[i+1:] {
			maxArea = max(p1.area(p2), maxArea)
		}
	}

	return maxArea
}

func part2(file string) uint {
	ps := parsePoints(file)
	ps.scale()

	maxX, maxY := uint(0), uint(0)
	for _, p := range ps {
		maxX = max(p.x+2, maxX)
		maxY = max(p.y+2, maxY)
	}

	g := newGrid(maxX, maxY)
	g.fill(&ps)

	maxArea := uint(0)
	for i, p1 := range ps {
		for _, p2 := range ps[i+1:] {
			area := p1.area(p2)
			if area <= maxArea {
				continue
			}
			if !g.contains(p1, p2) {
				continue	
			}
			maxArea = area
		}
	}
	return maxArea
}

func main() {
	fmt.Printf("part1(\"./inputs/day-09\"): %v\n", part1("./inputs/day-09"))
	fmt.Printf("part2(\"./inputs/day-09\"): %v\n", part2("./inputs/day-09"))
}
