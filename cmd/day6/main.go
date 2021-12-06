package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fish := make([]uint64, 9, 9)

	input, _ := os.ReadFile("internal/inputs/day6.txt")

	for _, n := range strings.Split(string(input), ",") {
		timer, err := strconv.Atoi(n)

		if err == nil {
			fish[timer]++
		}
	}

	for i := 0; i <= 255; i++ {
		temp := fish[1:9]
		temp = append(temp, fish[0])
		temp[6] += temp[8]
		fish = temp
	}

	var sum uint64 = 0
	for i := 0; i < 9; i++ {
		sum += fish[i]
	}

	fmt.Println(sum)
}
