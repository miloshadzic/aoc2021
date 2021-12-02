package main

import (
	"fmt"

	"github.com/miloshadzic/aoc2021/in"
	"github.com/miloshadzic/aoc2021/sub"
)

func main() {
	sub := sub.New()

	for _, command := range in.GetCommands("day2") {
		sub.Navigate(command)
	}

	fmt.Println(sub)
	fmt.Println(sub.Pos.H * sub.Pos.V)
}
