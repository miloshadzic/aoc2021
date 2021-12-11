package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/miloshadzic/aoc2021/geo"
)

type Cave struct {
	Step    uint8
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
			energy, err := strconv.ParseUint(string(r), 10, 8)

			if err == nil {
				cave.Energy[i][j] = int8(energy)
			}
		}
	}

	return cave
}

var adj [8]geo.P8 = [8]geo.P8{
	{X: 0, Y: 1},
	{X: 0, Y: -1},
	{X: -1, Y: 0},
	{X: 1, Y: 0},
	{X: 1, Y: 1},
	{X: 1, Y: -1},
	{X: -1, Y: 1},
	{X: -1, Y: -1},
}

func (cave *Cave) Next() {
	for i, row := range cave.Energy {
		for j := range row {
			cave.Flashed[i][j] = false
		}
	}

	for i, row := range cave.Energy {
		for j := range row {
			cave.Energy[i][j] = (cave.Energy[i][j] + 1) % 10
		}
	}

	for i, row := range cave.Energy {
		for j := range row {
			if cave.Energy[i][j] == 0 {
				cave.Flash(int8(i), int8(j))
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

		if x < 0 || x > int8(len(cave.Energy)-1) || y < 0 || y > int8(len(cave.Energy[0])-1) || cave.Flashed[x][y] ||
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
