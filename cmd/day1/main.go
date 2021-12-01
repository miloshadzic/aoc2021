package main

import (
	"fmt"
	"strconv"

	input "github.com/miloshadzic/aoc2020/internal/input"
)

func day1a() int64 {
	var prev, increases int64
	prev, increases = -1, 0

	for _, line := range input.GetLines("day1") {
		depth, _ := strconv.ParseInt(line, 10, 0)

		if depth > prev && prev > 0 {
			increases++
		}

		prev = depth
	}

	return increases
}

func day1b() int64 {
	var prev, increases int64

	depths := input.GetInt64s("day1")

	sum := depths[0] + depths[1] + depths[2]
	prev, increases = sum, 0

	for i := 3; i < len(depths); i++ {
		sum = sum + depths[i] - depths[i-3]

		if sum > prev {
			increases++
		}

		prev = sum
	}

	return increases
}

func main() {
	fmt.Println(day1a())
	fmt.Println(day1b())
}
