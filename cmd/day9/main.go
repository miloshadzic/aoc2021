package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/miloshadzic/aoc2021/geo"
)

type HeightMap struct {
	points [][]int
	basins [][]bool // actually a basin map

	lows     []geo.Point
	basinMap map[geo.Point]geo.Point // point -> low point
	sizeMap  map[geo.Point]int       // low point -> size of basin
}

var adj [4]geo.Point = [4]geo.Point{
	{X: 0, Y: 1},
	{X: 0, Y: -1},
	{X: -1, Y: 0},
	{X: 1, Y: 0},
}

func (hm *HeightMap) adjPoints(center geo.Point) []geo.Point {
	var points []geo.Point

	for _, p := range adj {
		x := center.X + p.X
		if x < 0 || x >= len(hm.points) {
			continue
		}

		y := center.Y + p.Y
		if y < 0 || y >= len(hm.points[0]) {
			continue
		}

		points = append(points, geo.Point{X: x, Y: y})
	}

	return points
}

func (hm *HeightMap) lowPoint(x, y int) bool {
	v, lower := hm.points[x][y], true

	for _, p := range hm.adjPoints(geo.Point{X: x, Y: y}) {
		lower = lower && hm.points[p.X][p.Y] > v
	}

	return lower
}

func (hm *HeightMap) riskLevel() int {
	var riskLevel int

	for _, p := range hm.lows {
		riskLevel += 1 + hm.points[p.X][p.Y]
	}

	return riskLevel
}

func (hm *HeightMap) scan(point, lowPoint geo.Point) {
	if _, ok := hm.basinMap[point]; !ok && hm.basins[point.X][point.Y] {
		hm.basinMap[point] = lowPoint
		hm.sizeMap[lowPoint]++

		for _, p := range hm.adjPoints(point) {
			hm.scan(p, lowPoint)
		}
	}
}

func main() {
	hm := HeightMap{}
	hm.basinMap = make(map[geo.Point]geo.Point)
	hm.sizeMap = make(map[geo.Point]int)

	f, _ := os.Open("internal/inputs/day9.txt")
	s := bufio.NewScanner(f)

	for s.Scan() {
		var pointRow []int
		var basinRow []bool

		for _, c := range s.Text() {
			n, err := strconv.Atoi(string(c))

			if err == nil {
				pointRow = append(pointRow, n)

				if n == 9 {
					basinRow = append(basinRow, false)
				} else {
					basinRow = append(basinRow, true)
				}
			}
		}
		hm.points = append(hm.points, pointRow)
		hm.basins = append(hm.basins, basinRow)
	}

	for i, row := range hm.points {
		for j := range row {
			if hm.lowPoint(i, j) {
				hm.lows = append(hm.lows, geo.Point{X: i, Y: j})
			}
		}
	}

	for _, lp := range hm.lows {
		hm.scan(lp, lp)
	}

	sizes := make([]int, 0, len(hm.sizeMap))
	for _, size := range hm.sizeMap {
		sizes = append(sizes, size)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	fmt.Println(hm.riskLevel())
	fmt.Println(sizes[0] * sizes[1] * sizes[2])
}
