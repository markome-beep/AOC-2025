package main

import (
	"fmt"
	"iter"
	"strconv"
	"strings"
)

func doubleNums() iter.Seq[int] {
	return func(yeild func(int) bool) {
		base := 10
		for {
			for mul := base / 10; mul < base; mul++ {
				if !yeild(mul * (base + 1)) {
					return
				}
			}
			base *= 10

		}
	}
}

func day02() {
	sum := 0
	for r := range readLines("./inputs/day-02", ",") {
		r := strings.Split(r, "-")
		left, err := strconv.Atoi(strings.TrimSpace(r[0]))
		if err != nil {
			fmt.Println("RIP Left")
		}
		right, err := strconv.Atoi(strings.TrimSpace(r[1]))
		if err != nil {
			fmt.Println("RIP Right")
		}

		for val := range doubleNums() {
			if val > right {
				break
			}
			if val >= left {
				sum += val
			}
		}
	}
	fmt.Println(sum)
}

func day02_p2() {
	sum := 0
	for r := range readLines("./inputs/day-02", ",") {
		r := strings.Split(r, "-")
		left, err := strconv.Atoi(strings.TrimSpace(r[0]))
		if err != nil {
			fmt.Println("RIP Left")
		}
		right, err := strconv.Atoi(strings.TrimSpace(r[1]))
		if err != nil {
			fmt.Println("RIP Right")
		}

		Num:
		for i := left; i <= right; i++ {
			num := strconv.Itoa(i)
			Split:
			for splits := 1; splits < len(num); splits++ {
				if len(num)%splits != 0 {
					continue
				}
				index := splits
				for index < len(num) {
					if num[:splits] != num[index:index+splits] {
						continue Split
					}
					index += splits
				}
				sum += i
				continue Num

			}
		}
	}
	fmt.Println(sum)
}
