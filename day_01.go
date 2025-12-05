package main

import (
	"fmt"
	"strconv"
)

func day01() {
	position := 50
	count := 0
	for line := range readLines("./inputs/day-01", "\n") {

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

func day01_p2() {
	position := 50
	count := 0
	for line := range readLines("./inputs/day-01", "\n") {

		value, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println("RIP")
			continue
		}

		count += value / 100

		value %= 100

		switch line[0] {
		case 'L':
			if position == 0 { count --}
			position -= value
		case 'R':
			position += value
		default:
			fmt.Println("RIP")
		}

		if position >= 100 || position <= 0 {
			count++
		}

		position = (position % 100 + 100) % 100

		fmt.Println("Position: ", position)
		fmt.Println("Count: ", count)
	}
	fmt.Println("Answer: ", count)
}
