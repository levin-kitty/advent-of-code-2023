package main

import (
	"fmt"
	"os"
	"strings"
)

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

	sum := 0
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			// not a number && continue
			if mat[i][j] < '0' || mat[i][j] > '9' {
				continue
			}

			// scan
			k := j
			isPart := false
			num := 0
			for k < len(mat[0]) && (mat[i][k] >= '0' && mat[i][k] <= '9') {
				num = 10*num + int(mat[i][k]-'0')
				if marked[i][k] {
					isPart = true
				}
				k++
			}

			// is part? then
			if isPart {
				sum += num
				fmt.Println(num)
			}
			j = k
		}
	}

	fmt.Println("answer:", sum)
}

func getMarked(mat []string) [][]bool {
	marked := make([][]bool, len(mat))
	for i := range marked {
		marked[i] = make([]bool, len(mat[0]))
	}

	for i := range mat {
		for j := range mat[0] {
			if isSymbol(mat[i][j]) {
				if j-1 >= 0 {
					marked[i][j-1] = true
				}
				if j+1 < len(mat[0]) {
					marked[i][j+1] = true
				}
				if i-1 >= 0 {
					marked[i-1][j] = true
					if j-1 >= 0 {
						marked[i-1][j-1] = true
					}
					if j+1 < len(mat[0]) {
						marked[i-1][j+1] = true
					}
				}
				if i+1 >= 0 {
					marked[i+1][j] = true
					if j-1 >= 0 {
						marked[i+1][j-1] = true
					}
					if j+1 < len(mat[0]) {
						marked[i+1][j+1] = true
					}
				}
			}
		}
	}

	return marked
}

func isSymbol(b byte) bool {
	return b != '.' && (b < '0' || b > '9')
}
