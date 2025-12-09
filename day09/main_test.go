package main

import "testing"

func Test_part1(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		file string
		want uint
	}{
		{"Example", "../inputs/day-09-test", 50},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := part1(tt.file)
			// TODO: update the condition below to compare got with tt.want.
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
		want uint
	}{
		{"Example", "../inputs/day-09-test", 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := part2(tt.file)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

