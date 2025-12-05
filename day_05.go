package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type bucket struct {
	left int
	right int
}

func parseDB(s string) []bucket {
	lines := strings.Split(s, "\n")
	db := make([]bucket, 0, len(lines))
	for _, line := range lines {
		r := strings.Split(line, "-")
		left, err := strconv.Atoi(strings.TrimSpace(r[0]))
		if err != nil {
			fmt.Println("RIP Left")
		}
		right, err := strconv.Atoi(strings.TrimSpace(r[1]))
		if err != nil {
			fmt.Println("RIP Right")
		}
		db = append(db, bucket{left, right})
	}
	return db
}

func day05()  {
	var tmp []string
	for part := range readLines("./inputs/day-05", "\n\n") {
		tmp = append(tmp, part)
	}

	db := parseDB(tmp[0])
	sum := 0

	Val:
	for _, val := range strings.Split(tmp[1], "\n") {
		num, err := strconv.Atoi(strings.TrimSpace(val))
		if err != nil {
			continue
		}
		for _, b := range db {
			if b.left <= num && num <= b.right {
				sum ++
				continue Val
			}
		}
	}
	fmt.Printf("sum: %v\n", sum)
}

func day05_p2()  {
	var tmp []string
	for part := range readLines("./inputs/day-05", "\n\n") {
		tmp = append(tmp, part)
	}

	db := parseDB(tmp[0])
	sort.Slice(db, func(i, j int) bool {
		return db[i].right-db[i].left > db[j].right-db[j].left
	})
	sum := 0

	Bucket:
	for i, buck := range db {
		for _, prev_buck := range db[:i] {
			if buck.left > prev_buck.right {
				continue
			} else if buck.right < prev_buck.left {
				continue
			} else if buck.right > prev_buck.right {
				buck.left = prev_buck.right + 1
			} else if buck.left < prev_buck.left { 
				buck.right = prev_buck.left - 1
			} else {
				continue Bucket
			}
		}
		sum += buck.right - buck.left + 1
	}
	fmt.Printf("sum: %v\n", sum)
}
