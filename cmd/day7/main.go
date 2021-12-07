package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/miloshadzic/aoc2021/util"
)

type fuelCost struct {
	costs []int
}

func InitFuelCost() fuelCost {
	return fuelCost{[]int{0}}
}

func (c *fuelCost) get(step int) int {
	for step > len(c.costs)-1 {
		c.costs = append(c.costs, step+c.get(step-1))
	}
	return c.costs[step]
}

func main() {
	var positions []int
	var fuel [1050][1050]int
	cost := InitFuelCost()

	input, _ := os.ReadFile("internal/inputs/day7.txt")

	for _, s := range strings.Split(string(input), ",") {
		n, err := strconv.Atoi(s)

		if err == nil {
			positions = append(positions, n)
		}
	}

	for i, p := range positions {
		for j := range fuel {
			if i == 0 {
				fuel[i][j] = cost.get(util.Abs(p - j))
			} else {
				fuel[i][j] = cost.get(util.Abs(p-j)) + fuel[i-1][j]
			}
		}
	}

	var minIndex, minFuel int = -1, 0
	for i, v := range fuel[len(positions)-1] {
		if v < minFuel || minIndex < 0 {
			minFuel, minIndex = v, i
		}
	}

	fmt.Println(minFuel)
}
