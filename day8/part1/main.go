package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	instructions, graph, idxMap := GetGraph("../input")

	// solve
	step := 0
	pos := 0
	cur := idxMap["AAA"]
	for cur != idxMap["ZZZ"] {
		switch instructions[pos] {
		case 'L':
			cur = graph[cur][0]
		case 'R':
			cur = graph[cur][1]
		}
		pos = (pos + 1) % len(instructions)
		step++
	}
	fmt.Println(step)
}

// instructions, graph, idxMap
func GetGraph(fileName string) (string, [][2]int, map[string]int) {
	lines := FileToLines(fileName)

	instructions := lines[0]

	idxMap := make(map[string]int)
	idx := 0
	for l := 2; l < len(lines); l++ {
		tokens := strings.Split(lines[l], " ")
		idxMap[tokens[0]] = idx
		idx++
	}

	graph := make([][2]int, len(idxMap))
	for l := 2; l < len(lines); l++ {
		tokens := strings.Split(lines[l], " ")
		graph[idxMap[tokens[0]]][0] = idxMap[tokens[2][1:len(tokens[2])-1]]
		graph[idxMap[tokens[0]]][1] = idxMap[tokens[3][:len(tokens[3])-1]]
	}

	return instructions, graph, idxMap
}

func FileToLines(fileName string) []string {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), "\n")
}
