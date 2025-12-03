package main

import (
	"fmt"
	"strconv"
)

func day01() {
	position := 50
	count := 0
	for line := range readLines("./inputs/day-01-test", "\n") {

		value, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println("RIP")
			continue
		}

		switch line[0] {
		case 'L':
			position += value
		case 'R':
			position -= value
		default:
			fmt.Println("RIP")
		}

			position %= 100
		if position == 0 {
			count += 1
		}
	}
	fmt.Println("Answer: ", count)
}
