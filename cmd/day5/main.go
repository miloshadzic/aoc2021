package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/miloshadzic/aoc2021/geo"
)

func main() {
	var grid [1000][1000]int
	var line geo.Line
	overlaps := 0

	f, err := os.Open("internal/inputs/day5.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	// Read the line segments
	for s.Scan() {
		pair := strings.Split(s.Text(), " -> ")

		A := strings.Split(pair[0], ",")
		B := strings.Split(pair[1], ",")

		aX, _ := strconv.Atoi(A[0])
		aY, _ := strconv.Atoi(A[1])
		bX, _ := strconv.Atoi(B[0])
		bY, _ := strconv.Atoi(B[1])

		line = geo.Line{
			A: geo.Point{X: aX, Y: aY},
			B: geo.Point{X: bX, Y: bY},
		}

		for i := 0; i < line.Len(); i++ {
			x := line.A.X + i*line.XStep()
			y := line.A.Y + i*line.YStep()

			grid[x][y]++

			if grid[x][y] == 2 {
				overlaps++
			}
		}
	}

	fmt.Println(overlaps)
}
