package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func day06() {
	input := make([][]string, 0)

	for line := range readLines("./inputs/day-06", "\n") {
		row := strings.Fields(line)
		input = append(input, row)
	}

	cols := make([][]*big.Int, len(input[0]))

	for _, row := range input[:len(input)-1] {
		for col, val := range row {
			num, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				fmt.Printf("val: %v\n", val)
				fmt.Println("RIP num")
				continue
			}
			cols[col] = append(cols[col], big.NewInt(num))
		}
	}

	total := big.NewInt(0)
	for col, op := range input[len(input)-1] {
		var val *big.Int
		switch op {
		case "+":
			sum := big.NewInt(0)
			for _, num := range cols[col] {
				sum.Add(sum, num)
			}
			val = sum
		case "*":
			prod := big.NewInt(1)
			for _, num := range cols[col] {
				prod.Mul(prod, num)
			}
			val = prod
		default:
			fmt.Printf("op: %v\n", op)
			fmt.Println("RIP OP")
		}
		total.Add(total, val)
	}
	fmt.Printf("total: %v\n", total)
}

func day06_p2() {
	input := make([]string, 0)

	for line := range readLines("./inputs/day-06", "\n") {
		input = append(input, line)
	}

	rot := make([][]rune, len(input[0]))
	for i := range rot {
		rot[i] = make([]rune, len(input[:len(input)-1]))
	}

	for i, row := range input[:len(input)-1] {
		for j, char := range row {
			rot[j][i] = char
		}
	}

	var curr rune
	tmpTotal := 0
	total := 0

	for i, op := range input[len(input)-1] {
		switch op {
		case '+':
			// fmt.Printf("tmpTotal: %v\n", tmpTotal)
			total += tmpTotal
			tmpTotal = 0
			curr = op
			// fmt.Printf("op: %v\n", string(op))
		case '*':
			// fmt.Printf("tmpTotal: %v\n", tmpTotal)
			total += tmpTotal
			tmpTotal = 1
			curr = op
			// fmt.Printf("op: %v\n", string(op))
		case ' ':
		default:
			fmt.Println("RIP")
		}
		num, err := strconv.Atoi(strings.TrimSpace(string(rot[i])))
		if err != nil {
			continue
		}
		// fmt.Printf("num: %v\n", num)
		switch curr {
		case '+':
			tmpTotal += num
		case '*':
			tmpTotal *= num
		default:
			fmt.Println("RIP")
		}
	}
	total += tmpTotal

	fmt.Printf("total: %v\n", total)
}
