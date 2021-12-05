package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// A line segment
type Line struct {
	A Point
	B Point
}

type Point struct {
	X int
	Y int
}

func main() {
	var grid [1000][1000]int
	var lines []Line

	// Initialize the grid
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = 0
		}
	}

	f, err := os.Open("internal/inputs/day5.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	// Read the line segments
	for s.Scan() {
		line := Line{Point{}, Point{}}

		pair := strings.Split(s.Text(), " -> ")

		A := strings.Split(pair[0], ",")
		B := strings.Split(pair[1], ",")

		line.A.X, _ = strconv.Atoi(A[0])
		line.A.Y, _ = strconv.Atoi(A[1])
		line.B.X, _ = strconv.Atoi(B[0])
		line.B.Y, _ = strconv.Atoi(B[1])

		lines = append(lines, line)
	}

	for _, line := range lines {
		for i := 0; i < line.len(); i++ {
			x := line.A.X + i*line.xStep()
			y := line.A.Y + i*line.yStep()

			grid[x][y]++
		}
	}

	overlaps := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] > 1 {
				overlaps++
			}
		}
	}

	fmt.Println(overlaps)
}

// Length of a line segment
func (line *Line) len() int {
	return 1 + Max(Abs(line.dx()), Abs(line.dy()))
}

func (line *Line) dx() int {
	return line.B.X - line.A.X
}

func (line *Line) dy() int {
	return line.B.Y - line.A.Y
}

func (line *Line) xStep() int {
	if line.dx() > 0 {
		return 1
	} else if line.dx() < 0 {
		return -1
	}

	return 0
}

func (line *Line) yStep() int {
	if line.dy() > 0 {
		return 1
	} else if line.dy() < 0 {
		return -1
	}

	return 0
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
