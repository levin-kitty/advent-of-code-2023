package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	Index  int
	Rounds []Round
}

type Round struct {
	Cubes map[string]int
}

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	games := parse(string(data))
	fmt.Println("answer:", solve(games))
}

func solve(games []Game) int {
	sum := 0
	for _, game := range games {

		minimumCubes := make(map[string]int)
		for _, round := range game.Rounds {
			for _, color := range []string{"red", "green", "blue"} {
				if round.Cubes[color] > minimumCubes[color] {
					minimumCubes[color] = round.Cubes[color]
				}
			}
		}

		power := 1
		for _, color := range []string{"red", "green", "blue"} {
			if minimumCubes[color] > 0 {
				power *= minimumCubes[color]
			}
		}
		sum += power
	}
	return sum
}

func parse(data string) (games []Game) {
	for _, line := range strings.Split(data, "\n") {
		tokens := strings.Split(line, ":")
		if len(tokens) != 2 {
			break
		}

		// get game index
		game, err := strconv.Atoi(strings.Split(tokens[0], " ")[1])
		if err != nil {
			panic(err)
		}
		g := Game{
			Index:  game,
			Rounds: []Round{},
		}

		rounds := strings.Split(tokens[1], ";")
		for _, round := range rounds {
			r := Round{
				Cubes: map[string]int{},
			}

			cubeSet := strings.Split(round, ",")
			for _, cubes := range cubeSet {
				cubes = strings.Trim(cubes, " ")
				cubesDetail := strings.Split(cubes, " ")
				r.Cubes[cubesDetail[1]], err = strconv.Atoi(cubesDetail[0])
				if err != nil {
					panic(err)
				}
			}
			g.Rounds = append(g.Rounds, r)
		}
		games = append(games, g)
	}
	return
}
