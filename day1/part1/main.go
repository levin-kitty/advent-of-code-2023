package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	sum := 0
	for _, line := range strings.Split(string(data), "\n") {
		for pos := 0; pos < len(line); pos++ {
			if isDigit(line[pos]) {
				sum += 10 * getDigit(line[pos])
				break
			}
		}
		for pos := len(line) - 1; pos >= 0; pos-- {
			if isDigit(line[pos]) {
				sum += getDigit(line[pos])
				break
			}
		}
	}

	fmt.Println("answer:", sum)
}

func isDigit(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	} else {
		return false
	}
}

func getDigit(c byte) int {
	return int(c - '0')
}
