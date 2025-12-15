package main

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"

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
	var count int32
	var numLines int32

	buff := 3
	results := make(chan uint, buff)
	jobs := make(chan string, 200)

	for range buff {
		go func() {
			for line := range jobs {
				m := NewMachine(line)
				// results <- m.Joltage_BFS()
				results <- m.Joltage()
				wg.Done()
				atomic.AddInt32(&count, -1)
			}
		}()
	}

	for line := range shared.ReadLines(file, "\n") {
		wg.Add(1)
		atomic.AddInt32(&count, 1)
		atomic.AddInt32(&numLines, 1)
		jobs <- line
	}

	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	var total uint = 0
	for r := range results {
		total += r
		log.Printf("%v / %v\n", count, numLines)
	}
	return total
}

func main() {
	fmt.Printf("part1: %v\n", part1("./inputs/day-10"))
	fmt.Printf("part2: %v\n", part2("./inputs/day-10"))
}
