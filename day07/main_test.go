package main

import "testing"

func Test_part_1(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named file parameters for target function.
		file string
		want  int
	}{
		{"Example", "../inputs/day-07-test", 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := part_1(tt.file)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("part_1() = %v, want %v", got, tt.want)
			}
		})
	}
}


func Test_part_2(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		file string
		want int
	}{
		{"Example", "../inputs/day-07-test", 40},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := part_2(tt.file)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("part_2() = %v, want %v", got, tt.want)
			}
		})
	}
}

