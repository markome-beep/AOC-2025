package main

import "fmt"

func maxi(r []int) (int, int) {
	if len(r) == 0 {
		return -1, 0
	}

	i := 0
	m := r[0]
	for j, v := range r {
		if v > m {
			i = j
			m = v
		}
	}
	return i, m
}

func day03() {
	sum := 0
	for line := range readLines("./inputs/day-03", "\n") {
		nums := make([]int, len(line))
		for i, char := range line {
			nums[i] = int(char - '0')
		}
		i, m := maxi(nums[:len(nums)-1])
		_, m2 := maxi(nums[i+1:])
		sum += m2 + m*10
	}
	fmt.Println(sum)
}

func day03_p2() {
	sum := 0
	for line := range readLines("./inputs/day-03", "\n") {
		nums := make([]int, len(line))
		for i, char := range line {
			nums[i] = int(char - '0')
		}
		
		var m, k int
		j := 0
		mul := 100_000_000_000
		for i := range(12) {
			k, m = maxi(nums[j:len(nums)-(11-i)])
			j += k + 1
			sum += m * mul
			mul /= 10
		}
	}
	fmt.Println(sum)
}
