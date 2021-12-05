package geo

import "github.com/miloshadzic/aoc2021/util"

// A line segment
type Line struct {
	A Point
	B Point
}

// Length of a line segment
func (line *Line) Len() int {
	return 1 + util.Max(util.Abs(line.dx()), util.Abs(line.dy()))
}

func (line *Line) dx() int {
	return line.B.X - line.A.X
}

func (line *Line) dy() int {
	return line.B.Y - line.A.Y
}

func (line *Line) XStep() int {
	if line.dx() > 0 {
		return 1
	} else if line.dx() < 0 {
		return -1
	}

	return 0
}

func (line *Line) YStep() int {
	if line.dy() > 0 {
		return 1
	} else if line.dy() < 0 {
		return -1
	}

	return 0
}
