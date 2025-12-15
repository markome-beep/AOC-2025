package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/markome-beep/AOC-2025/shared"
)

type present struct {
	hash_count int
	shape      [][]bool
}

type tree struct {
	shape         [2]int
	present_count []int
}

func parseInput(file string) ([]present, []tree) {
	state := 0
	ps := make([]present, 0)
	ts := make([]tree, 0)
	for line := range shared.ReadLines(file, "\n") {
		switch state {
		case 0:
			re := regexp.MustCompile(`\d:$`)
			if re.MatchString(line) {
				ps = append(ps, present{0, make([][]bool, 0)})
				state = 1
				continue
			}

			parts := strings.Split(line, " ")
			shape := strings.Split(parts[0][:len(parts[0])-1], "x")
			width, _ := strconv.Atoi(shape[0])
			depth, _ := strconv.Atoi(shape[1])
			p_count := make([]int, len(parts)-1)
			for i, p := range parts[1:] {
				num, _ := strconv.Atoi(p)
				p_count[i] = num
			}
			
			ts = append(ts, tree{[2]int{width, depth}, p_count})
		case 1:
			if strings.TrimSpace(line) == "" {
				state = 0
				continue
			}
			row := make([]bool, len(line))
			for i, r := range line {
				if r == '#' {
					ps[len(ps)-1].hash_count += 1
					row[i] = true
				}
			}
			ps[len(ps)-1].shape = append(ps[len(ps)-1].shape, row)
		}
	}
	return ps, ts
}

func part1(file string) int {
	ps, ts := parseInput(file)
	count := 0
	
	for _, t := range ts {
		total_area := t.shape[0] * t.shape[1]
		acc_area := 0
		for i, pc := range t.present_count {
			acc_area += ps[i].hash_count * pc
		}
		if acc_area <= total_area {
			count ++
		}
	}


	return count
}

func main() {
	fmt.Printf("part1(\"./inputs/day-12\"): %v\n", part1("./inputs/day-12"))
}
