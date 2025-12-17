package main

import (
	"fmt"
	"sync"

	"github.com/markome-beep/AOC-2025/shared"
)

func part1(file string) uint {
	var wg sync.WaitGroup

	buff := 100
	results := make(chan uint, buff)

	for line := range shared.ReadLines(file, "\n") {
		wg.Add(1)
		go func(line string) {
			defer wg.Done()
			m := NewMachine(line)
			results <- m.Indicator()
		}(line)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var total uint = 0
	for r := range results {
		total += r
	}
	return total
}

func part2(file string) uint {
	var wg sync.WaitGroup

	buff := 100
	results := make(chan uint, buff)

	for line := range shared.ReadLines(file, "\n") {
		wg.Add(1)
		go func(line string) {
			defer wg.Done()
			m := NewMachine(line)
			results <- m.Joltage()
		}(line)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var total uint = 0
	for r := range results {
		total += r
	}
	return total
}

func main() {
	fmt.Printf("part1: %v\n", part1("./inputs/day-10"))
	fmt.Printf("part2: %v\n", part2("./inputs/day-10"))
}
