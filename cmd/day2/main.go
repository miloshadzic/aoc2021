package main

import (
	"fmt"
	"strconv"

	"github.com/miloshadzic/aoc2021/internal/input"
)

type Submarine struct {
	aim  int64
	hPos int64
	vPos int64
}

func main() {
	sub := Submarine{0, 0, 0}

	course := input.GetCommands("day2")

	for _, command := range course {
		amount, _ := strconv.ParseInt(command[1], 10, 0)

		switch command[0] {
		case "forward":
			sub.hPos += amount
			sub.vPos += amount * sub.aim
		case "up":
			sub.aim -= amount
		case "down":
			sub.aim += amount
		}
	}

	fmt.Println(sub)
	fmt.Println(sub.hPos * sub.vPos)
}
