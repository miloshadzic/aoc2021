package main

import (
	"fmt"

	"github.com/miloshadzic/aoc2021/bingo"
)

func main() {
	game := bingo.FromInput("day4")
	game.Run()
	fmt.Println(game.Wins[game.Order[len(game.Wins)-1]])
}
