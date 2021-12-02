package main

import (
	"fmt"

	"github.com/miloshadzic/aoc2021/internal/input"
)

type Submarine struct {
	aim  int64
	hPos int64
	vPos int64
}

func (sub *Submarine) Navigate(cmd input.Command) {
	switch cmd.Direction {
	case input.Forward:
		sub.hPos += cmd.Amount
		sub.vPos += cmd.Amount * sub.aim
	case input.Up:
		sub.aim -= cmd.Amount
	case input.Down:
		sub.aim += cmd.Amount
	}
}

func main() {
	sub := Submarine{0, 0, 0}

	for _, command := range input.GetCommands("day2") {
		sub.Navigate(command)
	}

	fmt.Println(sub)
	fmt.Println(sub.hPos * sub.vPos)
}
