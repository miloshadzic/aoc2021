package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/miloshadzic/aoc2021/geo"
)

type Cave struct {
	Energy  [10][10]int8
	Flashed [10][10]bool
	Flashes uint
}

func InitCave(input string) Cave {
	f, _ := os.Open(fmt.Sprintf("internal/inputs/%s.txt", input))
	s := bufio.NewScanner(f)

	cave := Cave{}

	for i := 0; i < 10; i++ {
		s.Scan()

		for j, r := range s.Text() {
			cave.Energy[i][j] = int8(r - '0')
		}
	}

	return cave
}

var adj [8]geo.P8 = [8]geo.P8{
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
	var i, j int8 // reusing this for all the iterations

	for i = 0; i < 10; i++ {
		for j = 0; j < 10; j++ {
			cave.Flashed[i][j] = false
			cave.Energy[i][j] = (cave.Energy[i][j] + 1) % 10
		}
	}

	for i = 0; i < 10; i++ {
		for j = 0; j < 10; j++ {
			if cave.Energy[i][j] == 0 {
				cave.Flash(i, j)
			}
		}
	}
}

func (cave *Cave) Flash(i, j int8) {
	if cave.Flashed[i][j] {
		return
	}

	cave.Flashes++

	for _, p := range adj {
		x := i + p.X
		y := j + p.Y

		if x < 0 || x > 9 || y < 0 || y > 9 || cave.Flashed[x][y] ||
			cave.Energy[x][y] == 0 {
			continue
		} else {
			cave.Energy[x][y] = (cave.Energy[x][y] + 1) % 10

			if cave.Energy[x][y] == 0 {
				cave.Flash(x, y)
			}
		}
	}

	cave.Flashed[i][j] = true
}

func (cave *Cave) AllFlash() bool {
	all := true

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			all = all && cave.Flashed[i][j]

			if !all {
				return false
			}
		}
	}

	return all
}

func main() {
	cave := InitCave("day11")

	for i := 0; i < 99999999; i++ {
		if i == 100 {
			fmt.Println(cave.Flashes)
		}

		cave.Next()

		if cave.AllFlash() {
			fmt.Println("Step", i+1)
			break
		}
	}
}
