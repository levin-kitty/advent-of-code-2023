package main

import "os"

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}
}
