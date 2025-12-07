package main

import (
	"fmt"

	"github.com/markome-beep/AOC-2025/shared"
)

func makeChart(file string) [][]rune {
	chart := make([][]rune, 0)
	for row := range shared.ReadLines(file, "\n") {
		r := make([]rune, len(row))
		for i, char := range row {
			r[i] = char
		}
		r = append([]rune{'.'}, r...)
		r = append(r, '.')
		chart = append(chart, r)
	}
	padding := make([]rune, len(chart[0]))
	for i := range padding {
		padding[i] = '.'
	}

	chart = append([][]rune{padding}, chart...)
	chart = append(chart, padding)
	return chart
}

func part_1() {
	neighbors := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	chart := makeChart("./inputs/day-04")
	sum := 0

	for x, row := range chart {
		if x == 0 || x == len(chart)-1 {
			continue
		}
		for y := range row {
			if y == 0 || y == len(chart[x])-1 {
				continue
			}

			if chart[x][y] == '.' {
				continue
			}
			buddies := 0
			for _, offset := range neighbors {
				if chart[x+offset[0]][y+offset[1]] != '.' {
					buddies++
				}
			}
			if buddies < 4 {
				sum++
			}
		}
	}
	fmt.Printf("sum: %v\n", sum)

}

func part_2() {
	neighbors := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	chart := makeChart("./inputs/day-04")
	sum := 0
	found := true

	for found {
		found = false
		for x, row := range chart {
			if x == 0 || x == len(chart)-1 {
				continue
			}
			for y := range row {
				if y == 0 || y == len(chart[x])-1 {
					continue
				}

				if chart[x][y] == '.' {
					continue
				}
				buddies := 0
				for _, offset := range neighbors {
					if chart[x+offset[0]][y+offset[1]] != '.' {
						buddies++
					}
				}
				if buddies < 4 {
					chart[x][y] = '.'
					sum++
				}
			}
		}
	}
	fmt.Printf("sum: %v\n", sum)
}

func main() {
	part_1()
	part_2()
}
