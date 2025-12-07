package main

import (
	"fmt"

	"github.com/markome-beep/AOC-2025/shared"
)

func part_1(file string) int {
	var lasers []bool
	splits := 0
	for line := range shared.ReadLines(file, "\n") {

		for j, char := range line {
			switch char {
			case 'S':
				lasers = make([]bool, len(line))
				lasers[j] = true
			case '^':
				if !lasers[j] {
					continue
				}
				lasers[j] = false
				lasers[j-1] = true
				lasers[j+1] = true
				splits++
			case '.':
			default:
				fmt.Println("RIP")
			}
		}
	}

	fmt.Printf("splits: %v\n", splits)
	return splits
}

func part_2(file string) int {
	timelines := 0

	var lasers []int
	for line := range shared.ReadLines(file, "\n") {

		for j, char := range line {
			switch char {
			case 'S':
				lasers = make([]int, len(line))
				lasers[j]++
			case '^':
				if lasers[j] == 0 {
					continue
				}
				lasers[j-1] += lasers[j]
				lasers[j+1] += lasers[j]
				lasers[j] = 0

			case '.':
			default:
				fmt.Println("RIP")
			}
		}
	}

	for _, v := range lasers {
		timelines += v
	}

	fmt.Printf("timelines: %v\n", timelines)
	return timelines
}

func main() {
	part_1("./inputs/day-07")
	part_2("./inputs/day-07")
}
