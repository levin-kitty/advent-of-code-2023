package main

import (
	"fmt"
	"os"
	"strings"
)

var digitString = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	sum := 0
	for _, line := range strings.Split(string(data), "\n") {
		fmt.Println(line)
		for pos := 0; pos < len(line); pos++ {
			digit, ok := getDigit(line, pos)
			if ok {
				sum += 10 * digit
				break
			}
		}

		for pos := len(line) - 1; pos >= 0; pos-- {
			digit, ok := getDigit(line, pos)
			if ok {
				sum += digit
				break
			}
		}
	}

	fmt.Println("answer:", sum)
}

func getDigit(line string, pos int) (int, bool) {
	if line[pos] >= '0' && line[pos] <= '9' {
		return int(line[pos] - '0'), true
	}

	for i, d := range digitString {
		if pos+len(d) <= len(line) && strings.HasPrefix(line[pos:], d) {
			return i, true
		}
	}

	return 0, false
}
