package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Map 표현은

func main() {
	lines := fileToLines("../input")
	almanac := parse(lines)
	fmt.Println("answer:", almanac.Solve())

	//_, light := almanac.MapByKeySrc("water")
	//fmt.Printf("%+v\n", light)
	//fmt.Println(RangesList(light).Dest(81))
}

type Key struct {
	Src string
	Dst string
}

type Ranges struct {
	SrcStart int
	Added    int
	Len      int
}

type RangesList []Ranges

func (r RangesList) Len() int {
	return len(r)
}

func (r RangesList) Less(i, j int) bool {
	return r[i].SrcStart < r[j].SrcStart
}

func (r RangesList) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

type Almanac struct {
	Seeds []int
	Maps  map[Key][]Ranges
}

func (a Almanac) Solve() int {
	minLocation := 1_000_000_000
	for _, seed := range a.Seeds {
		keysrc := "seed"

		next := seed
		for keysrc != "location" {
			key, rl := a.MapByKeySrc(keysrc)
			keysrc = key.Dst
			next = RangesList(rl).Dest(next)
		}
		if next < minLocation {
			minLocation = next
		}
	}
	return minLocation
}

// 1) Find Key by Key.Src
func (a Almanac) MapByKeySrc(src string) (Key, []Ranges) {
	for k, m := range a.Maps {
		if k.Src == src {
			return k, m
		}
	}
	panic("unreachable code: src=" + src)
}

// 2) Find Range
func (rl RangesList) Dest(src int) int {
	for i := 0; i < len(rl); i++ {
		if src >= rl[i].SrcStart && src < rl[i].SrcStart+rl[i].Len {
			return src + rl[i].Added
		}
	}
	return src
}

func parse(lines []string) Almanac {
	var key Key
	seeds := make([]int, 0)
	maps := make(map[Key][]Ranges)

	// 1) seeds
	for i := 0; i < len(lines); i++ {
		switch {
		case strings.Contains(lines[i], "seeds:"):
			seedStrs := strings.Split(lines[i], " ")
			for j := 1; j < len(seedStrs); j++ {
				s, err := strconv.Atoi(seedStrs[j])
				if err != nil {
					panic(err)
				}
				seeds = append(seeds, s)
			}

		case strings.Contains(lines[i], "map:"):
			tokens := strings.Split(strings.Split(lines[i], " ")[0], "-")
			key = Key{
				Src: tokens[0],
				Dst: tokens[2],
			}
			maps[key] = []Ranges{}

		case len(lines[i]) > 0:
			nums := strings.Split(lines[i], " ")
			destStart, err := strconv.Atoi(nums[0])
			if err != nil {
				panic(err)
			}
			sourceStart, err := strconv.Atoi(nums[1])
			if err != nil {
				panic(err)
			}
			rangeLen, err := strconv.Atoi(nums[2])
			if err != nil {
				panic(err)
			}
			maps[key] = append(maps[key], Ranges{
				SrcStart: sourceStart,
				Added:    destStart - sourceStart,
				Len:      rangeLen,
			})
		}
	}

	for _, rangesList := range maps {
		sort.Sort(RangesList(rangesList))
	}

	return Almanac{
		Seeds: seeds,
		Maps:  maps,
	}
}

func fileToLines(fileName string) []string {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), "\n")
}
