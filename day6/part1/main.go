package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines := FileToLines("../input")
	/*
		times, distances := Parse(lines)

		answer := 1
		for i := 0; i < len(times); i++ {
			answer *= NumberOfWays(times[i], distances[i])
			fmt.Println(NumberOfWays(times[i], distances[i]))
		}
		fmt.Println(answer)
	*/

	time, distance := ParsePartTwo(lines)
	fmt.Println(time, distance)
	fmt.Println(NumberOfWays(time, distance))
}

func NumberOfWays(t, d int64) int64 {
	// left:  (t - sqrt(t^2 - 4*d)) / 2
	// right: (t + sqrt(t^2 - 4*d)) / 2
	left, right := (float64(t)-math.Sqrt(float64(t*t-4*d)))/float64(2), (float64(t)+math.Sqrt(float64(t*t-4*d)))/float64(2)
	leftInt, rightInt := int64(math.Ceil(left)), int64(math.Floor(right))

	if (t-leftInt)*leftInt == d {
		leftInt++
	}
	if (t-rightInt)*rightInt == d {
		rightInt--
	}

	ways := rightInt - leftInt + 1
	if ways < 0 {
		panic("??")
	}
	return ways
}

func ParsePartTwo(lines []string) (time, distance int64) {
	space := regexp.MustCompile(`\s+`)
	time, _ = strconv.ParseInt(strings.Split(space.ReplaceAllString(lines[0], ""), ":")[1], 10, 0)
	distance, _ = strconv.ParseInt(strings.Split(space.ReplaceAllString(lines[1], ""), ":")[1], 10, 0)
	return
}

func Parse(lines []string) (times []int, distances []int) {
	space := regexp.MustCompile(`\s+`)
	for _, n := range strings.Split(space.ReplaceAllString(lines[0], " "), " ")[1:] {
		t, _ := strconv.Atoi(n)
		times = append(times, t)
	}
	for _, n := range strings.Split(space.ReplaceAllString(lines[1], " "), " ")[1:] {
		d, _ := strconv.Atoi(n)
		distances = append(distances, d)
	}
	return
}

func FileToLines(fileName string) []string {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), "\n")
}
