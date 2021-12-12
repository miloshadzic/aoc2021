package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/miloshadzic/aoc2021/geo"
)

func main() {
	cave := InitCave("day11")

	for i := 0; i < 99999999; i++ {
		if i == 100 {
			fmt.Println(cave.Flashes)
		}

		if cave.AllFlash() {
			fmt.Println("Step", i)
			break
		}

		cave.Next()
	}
}

type Octopus struct {
	Energy  int
	Flashed bool
}

func (o *Octopus) Boost() {
	o.Energy = (o.Energy + 1) % 10
}

type Cave struct {
	Octopi  [10][10]Octopus
	Flashes uint
}

func InitCave(input string) Cave {
	f, _ := os.Open(fmt.Sprintf("internal/inputs/%s.txt", input))
	s := bufio.NewScanner(f)

	cave := Cave{}

	for i := 0; i < 10; i++ {
		s.Scan()

		for j, r := range s.Text() {
			cave.Octopi[i][j] = Octopus{Energy: int(r - '0'), Flashed: false}
		}
	}

	return cave
}

var adj [8]geo.Point = [8]geo.Point{
	{X: 0, Y: 1},
	{X: 0, Y: -1},
	{X: 1, Y: 0},
	{X: 1, Y: 1},
	{X: 1, Y: -1},
	{X: -1, Y: 0},
	{X: -1, Y: 1},
	{X: -1, Y: -1},
}

func (cave *Cave) Next() {
	var i, j int // reusing this for all the iterations

	for i = 0; i < 10; i++ {
		for j = 0; j < 10; j++ {
			cave.Octopi[i][j].Flashed = false
			cave.Octopi[i][j].Boost()
		}
	}

	for i = 0; i < 10; i++ {
		for j = 0; j < 10; j++ {
			if cave.Octopi[i][j].Energy == 0 {
				cave.Flash(i, j)
			}
		}
	}
}

func (cave *Cave) Flash(i, j int) {
	if cave.Octopi[i][j].Flashed {
		return
	}

	var x, y int
	for _, p := range adj {
		x = i + p.X
		if x < 0 || x > 9 {
			continue
		}

		y = j + p.Y
		if y < 0 || y > 9 {
			continue
		}

		if x < 0 || x > 9 || y < 0 || y > 9 || cave.Octopi[x][y].Flashed ||
			cave.Octopi[x][y].Energy == 0 {
			continue
		} else {
			cave.Octopi[x][y].Boost()

			if cave.Octopi[x][y].Energy == 0 {
				cave.Flash(x, y)
			}
		}
	}

	cave.Flashes++
	cave.Octopi[i][j].Flashed = true
}

func (cave *Cave) AllFlash() bool {
	all := true

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			all = all && cave.Octopi[i][j].Flashed
		}
	}

	return all
}
