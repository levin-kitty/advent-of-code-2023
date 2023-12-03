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

		valid := 0
		for _, round := range game.Rounds {
			nKeys := len(round.Cubes)
			if cubes, ok := round.Cubes["red"]; ok {
				if cubes > 12 {
					break
				}
				nKeys--
			}
			if cubes, ok := round.Cubes["green"]; ok {
				if cubes > 13 {
					break
				}
				nKeys--
			}
			if cubes, ok := round.Cubes["blue"]; ok {
				if cubes > 14 {
					break
				}
				nKeys--
			}
			if nKeys > 0 {
				break
			} else {
				valid++
			}
		}

		if valid == len(game.Rounds) {
			sum += game.Index
		}
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
