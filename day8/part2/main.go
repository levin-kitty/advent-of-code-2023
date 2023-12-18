package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	instructions, graph, idxMap := GetGraph("../input")

	// currents
	curs := make([]int, 0)
	for k, v := range idxMap {
		if k[2] == 'A' {
			curs = append(curs, v)
		}
	}

	// ends
	ends := make(map[int]int, 0)
	for k, v := range idxMap {
		if k[2] == 'Z' {
			ends[v] = v
		}
	}

	fmt.Println(curs, ends)

	// move
	step := 1

	for pos := 0; ; pos = (pos + 1) % len(instructions) {

		matched := true
		for c := 0; c < len(curs); c++ {
			switch instructions[pos] {
			case 'L':
				curs[c] = graph[curs[c]][0]
			case 'R':
				curs[c] = graph[curs[c]][1]
			}

			if _, exists := ends[curs[c]]; !exists {
				matched = false
			}
		}

		if matched {
			break
		}
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
