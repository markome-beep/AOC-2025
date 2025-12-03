package main

import (
	"bufio"
	"fmt"
	"iter"
	"os"
	"strings"
)

func main() {
	day01()
}

func readLines(path string, split string) iter.Seq[string] {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return func(func(string) bool) {
			return
		}
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(customSplit(split))

	return func(yield func(string) bool) {
		defer file.Close() // Ensure the file is closed

		for scanner.Scan() { // Loop through each line
			line := scanner.Text()
			if !yield(line) {
				break
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file: ", err)
		}
	}
}

func customSplit(split string) func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {

		// Return nothing if at end of file and no data passed
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		// Find the index of the input of a newline followed by a
		// pound sign.
		if i := strings.Index(string(data), split); i >= 0 {
			return i + 1, data[0:i], nil
		}

		// If at end of file with data return the data
		if atEOF {
			return len(data), data, nil
		}

		return
	}
}
