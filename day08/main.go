package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/markome-beep/AOC-2025/shared"
)

type point struct {
	index int
	x     int
	y     int
	z     int
}

type pair struct {
	p1    point
	p2    point
	dist2 int
}

func (p1 point) dist2(p2 point) int {
	return (p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y) + (p1.z-p2.z)*(p1.z-p2.z)
}

func parsePoints(file string) []point {
	points := make([]point, 0)
	index := 0
	for line := range shared.ReadLines(file, "\n") {
		vals := strings.Split(line, ",")
		x, err := strconv.Atoi(vals[0])
		if err != nil {
			fmt.Println("RIP")
		}

		y, err := strconv.Atoi(vals[1])
		if err != nil {
			fmt.Println("RIP")
		}

		z, err := strconv.Atoi(vals[2])
		if err != nil {
			fmt.Println("RIP")
		}

		points = append(points, point{index, x, y, z})
		index++
	}
	return points
}

func distances(p []point) []pair {
	vals := make([]pair, 0)
	for i, p1 := range p {
		for _, p2 := range p[i+1:] {
			vals = append(vals, pair{p1, p2, p1.dist2(p2)})
		}
	}
	sort.Slice(vals, func(i, j int) bool {
		return vals[i].dist2 < vals[j].dist2
	})
	return vals
}

func part_1(file string, pairs int) int {
	points := parsePoints(file)
	connected := make([]int, len(points))
	for i := range connected {
		connected[i] = i
	}
	dists := distances(points)

	for conn := range pairs {
		if connected[dists[conn].p1.index] == connected[dists[conn].p2.index] {
			continue
		}

		update := connected[dists[conn].p2.index]
		for i := range connected {
			if connected[i] == update {
				connected[i] = connected[dists[conn].p1.index]
			}
		}
	}

	circuits := make([]int, len(points))
	for _, conn := range connected {
		circuits[conn]++
	}
	sort.Slice(circuits, func(i, j int) bool {
		return circuits[i] > circuits[j]
	})

	return circuits[0] * circuits[1] * circuits[2]
}

func part_2(file string) int {
	points := parsePoints(file)
	connected := make([]int, len(points))
	for i := range connected {
		connected[i] = i
	}
	dists := distances(points)

Conn:
	for _, conn := range dists {
		if connected[conn.p1.index] == connected[conn.p2.index] {
			continue
		}

		update := connected[conn.p2.index]
		for i := range connected {
			if connected[i] == update {
				connected[i] = connected[conn.p1.index]
			}
		}

		val := connected[0]
		for _, v := range connected[1:] {
			if v != val {
				continue Conn
			}
		}

		return conn.p1.x * conn.p2.x
	}

	return -1
}

func main() {
	fmt.Printf("part_1(\"./inputs/day-08\", 1000): %v\n", part_1("./inputs/day-08", 1000))

	fmt.Printf("part_2(\"./inputs/day-08\"): %v\n", part_2("./inputs/day-08"))
}
