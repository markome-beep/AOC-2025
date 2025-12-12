package main

import "testing"

func Test_part1(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		file string
		want int
	}{
		{"Example", "../inputs/day-11-test", 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := part1(tt.file)
			if got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}


func Test_part2(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		file string
		want int
	}{
		{"Example", "../inputs/day-11-test-2", 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := part2(tt.file)
			if got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

