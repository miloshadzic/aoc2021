package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var fish []uint8
	input, _ := os.ReadFile("internal/inputs/day6.txt")

	for _, n := range strings.Split(string(input), ",") {
		timer, err := strconv.Atoi(n)
		if err == nil {
			fish = append(fish, uint8(timer))
		}
	}

	fmt.Println(len(fish))

	for j := 0; j <= 255; j++ {
		for i := len(fish) - 1; i >= 0; i-- {
			if fish[i] == 0 {
				fish = append(fish, 8)
				fish[i] = 6
			} else {
				fish[i]--
			}
		}
	}

	fmt.Println(len(fish))
}
