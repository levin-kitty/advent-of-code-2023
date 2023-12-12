package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Key struct {
	Src string
	Dst string
}

type Function struct {
	X []int
	Y []int
}

type Almanac struct {
	Seeds []int
	Maps  map[Key]Function
}

func main() {
	almanac := ParseAlmanac(FileToLines("../input"))

	finalFunc := almanac.GetFuncBySrc("seed")
	nextKey := almanac.GetNextKey("seed")
	for nextKey != "location" {
		fmt.Println("next:", nextKey)
		finalFunc = MergeFunc(finalFunc, almanac.GetFuncBySrc(nextKey))
		nextKey = almanac.GetNextKey(nextKey)
	}

	fmt.Println(finalFunc.X)
	fmt.Println(finalFunc.Y)
	fmt.Println()

	min := 1_000_000_000
	//	for _, seed := range almanac.Seeds {
	//		min = Min(min, finalFunc.Min(seed, seed+1))
	//	}

	for i := 0; i < len(almanac.Seeds); i += 2 {
		min = Min(min, finalFunc.Min(almanac.Seeds[i], almanac.Seeds[i+1]))
	}

	fmt.Println(min)
}

/* Function */

func (f Function) Find(src int) int {
	for i := 1; i < len(f.X); i++ {
		if src < f.X[i] {
			return src + f.Y[i-1]
		}
	}
	return src
}

func (f Function) Min(src, dst int) int {
	// TODO: seed range to min location
	min := 1_000_000_000
	//fmt.Println("seeds", src, dst)

out:
	for i := 0; i < len(f.X)-1; i++ {
		switch {
		case f.X[i+1] < src:
			continue
		case f.X[i+1] < dst:
			if src+f.Y[i] < min {
				min = src + f.Y[i]
			}
			src = f.X[i+1]
		case f.X[i+1] >= dst:
			//fmt.Println(src, dst, f.Y[i], min)
			if src+f.Y[i] < min {
				min = src + f.Y[i]
			}
			src = dst
			//fmt.Println(min)
			break out
		}
	}
	if src < dst {
		if src < min {
			min = src
		}
	}
	return min
}

/* Transformation */

func MergeFunc(srcFunc Function, dstFunc Function) Function {

	mergedMap, mergedDiff := []int{}, []int{}

	// srcFunc 의 각 구간에 대해서 수행
	for i := 0; i < len(srcFunc.X)-1; i++ {
		fmt.Println("SrcFunx:", srcFunc.X[i], ",", srcFunc.X[i+1])

		// step 1: 구간의 Left, Right를 Y 만큼 더 함
		step1 := []int{srcFunc.X[i] + srcFunc.Y[i], srcFunc.X[i+1] + srcFunc.Y[i]}
		fmt.Println("Step1:", step1)

		// step 2: Left, Right 구간을 dstFunc을 검색 후 필요하면 분할
		step2 := []int{step1[0]}
		step2Diff := []int{}

	out:
		for j := 0; j < len(dstFunc.X)-1; j++ {
			switch {
			case dstFunc.X[j+1] <= step1[0]:
				continue
			case dstFunc.X[j+1] < step1[1]:
				step2 = append(step2, dstFunc.X[j+1])
				step2Diff = append(step2Diff, srcFunc.Y[i]+dstFunc.Y[j])
				step1[0] = dstFunc.X[j+1]
			case dstFunc.X[j+1] > step1[1]:
				step2 = append(step2, step1[1])
				step2Diff = append(step2Diff, srcFunc.Y[i]+dstFunc.Y[j])
				step1[0] = step1[1]
				break out
			}
		}
		if step1[0] < step1[1] {
			step2 = append(step2, step1[1])
			step2Diff = append(step2Diff, srcFunc.Y[i])
		}
		fmt.Println("Step2:", step2)
		fmt.Println("Step2 Diff:", step2Diff)

		// step 3: re transformation
		for j := 0; j < len(step2); j++ {
			step2[j] -= srcFunc.Y[i]
		}
		fmt.Println("Step3:", step2)
		fmt.Println("")

		// step 4: merge
		if len(mergedMap) == 0 {
			mergedMap = step2
			mergedDiff = step2Diff
		} else {
			mergedMap = append(mergedMap, step2[1:]...)
			mergedDiff = append(mergedDiff, step2Diff...)
		}
	}

	mergedDiff = append(mergedDiff, 0)
	return Function{
		X: mergedMap,
		Y: mergedDiff,
	}
}

/* Almanac */

func (a Almanac) GetFuncBySrc(str string) Function {
	for k, v := range a.Maps {
		if k.Src == str {
			return v
		}
	}
	return Function{}
}

func (a Almanac) GetNextKey(str string) string {
	for k, _ := range a.Maps {
		if k.Src == str {
			return k.Dst
		}
	}
	return ""
}

/* Parse Almanac */

type Range struct {
	Left  int
	Right int
	Diff  int
}

type Ranges []Range

func (r Ranges) Len() int {
	return len(r)
}

func (r Ranges) Less(i, j int) bool {
	return r[i].Left < r[j].Left
}

func (r Ranges) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func ParseAlmanac(lines []string) Almanac {
	seeds := make([]int, 0)
	maps := make(map[Key]Function)

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

			// Key
			key := Key{
				Src: tokens[0],
				Dst: tokens[2],
			}

			// Ranges
			r := make(Ranges, 0)
			for i = i + 1; i < len(lines) && len(lines[i]) > 0; i++ {
				nums := strings.Split(lines[i], " ")
				r = append(r, Range{
					Left:  MustAtoi(nums[1]),
					Right: MustAtoi(nums[1]) + MustAtoi(nums[2]),
					Diff:  MustAtoi(nums[0]) - MustAtoi(nums[1]),
				})
			}

			// Ranges to Slice
			sort.Sort(r)

			x, y := []int{-1}, []int{0}
			for i := 0; i < len(r); i++ {
				x, y = append(x, r[i].Left), append(y, r[i].Diff)
				if i+1 >= len(r) || r[i].Right < r[i+1].Left {
					x, y = append(x, r[i].Right), append(y, 0)
				}
			}

			maps[key] = Function{
				X: x,
				Y: y,
			}
		}
	}

	return Almanac{
		Seeds: seeds,
		Maps:  maps,
	}
}

func FileToLines(fileName string) []string {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), "\n")
}

func MustAtoi(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return n
}

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
