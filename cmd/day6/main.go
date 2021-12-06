package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var fish [9]uint64
	var temp uint64

	input, _ := os.ReadFile("internal/inputs/day6.txt")

	for _, n := range strings.Split(string(input), ",") {
		timer, err := strconv.Atoi(n)

		if err == nil {
			fish[timer]++
		}
	}

	for i := 0; i <= 255; i++ {
		fish[0], temp = temp, fish[0]

		fish[1], fish[0] = fish[0], fish[1]
		fish[2], fish[1] = fish[1], fish[2]
		fish[3], fish[2] = fish[2], fish[3]
		fish[4], fish[3] = fish[3], fish[4]
		fish[5], fish[4] = fish[4], fish[5]
		fish[6], fish[5] = fish[5], fish[6]
		fish[7], fish[6] = fish[6], fish[7]
		fish[8], fish[7] = fish[7], fish[8]

		fish[8] = temp
		fish[6] += temp
	}

	var sum uint64 = 0
	for i := 0; i < 9; i++ {
		sum += fish[i]
	}

	fmt.Println(sum)
}
