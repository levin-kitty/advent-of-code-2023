package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	I int
	J int
}

func main() {
	data, err := os.ReadFile("../input")
	if err != nil {
		panic(err)
	}

	mat := make([]string, 0)
	for _, line := range strings.Split(string(data), "\n") {
		mat = append(mat, line)
	}

	marked := getMarked(mat)

	gears := make(map[*Point][]int)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			// not a number && continue
			if mat[i][j] < '0' || mat[i][j] > '9' {
				continue
			}

			// scan
			k := j
			isPart := -1
			num := 0
			for k < len(mat[0]) && (mat[i][k] >= '0' && mat[i][k] <= '9') {
				num = 10*num + int(mat[i][k]-'0')
				if marked[i][k] != nil {
					isPart = k
				}
				k++
			}

			// is part? then
			if isPart >= 0 {
				if _, exists := gears[marked[i][isPart]]; exists {
					gears[marked[i][isPart]] = append(gears[marked[i][isPart]], num)
				} else {
					gears[marked[i][isPart]] = []int{num}
				}
			}
			j = k
		}
	}
	fmt.Println(gears)

	sum := 0
	for _, g := range gears {
		if len(g) == 2 {
			sum += (g[0] * g[1])
		}
	}
	fmt.Println(sum)
}

func getMarked(mat []string) [][](*Point) {
	marked := make([][](*Point), len(mat))
	for i := range marked {
		marked[i] = make([](*Point), len(mat[0]))
	}

	for i := range mat {
		for j := range mat[0] {
			if isSymbol(mat[i][j]) {
				p := Point{I: i, J: j}

				if j-1 >= 0 {
					marked[i][j-1] = &p
				}
				if j+1 < len(mat[0]) {
					marked[i][j+1] = &p
				}
				if i-1 >= 0 {
					marked[i-1][j] = &p
					if j-1 >= 0 {
						marked[i-1][j-1] = &p
					}
					if j+1 < len(mat[0]) {
						marked[i-1][j+1] = &p
					}
				}
				if i+1 >= 0 {
					marked[i+1][j] = &p
					if j-1 >= 0 {
						marked[i+1][j-1] = &p
					}
					if j+1 < len(mat[0]) {
						marked[i+1][j+1] = &p
					}
				}
			}
		}
	}

	return marked
}

func isSymbol(b byte) bool {
	return b == '*'
}
