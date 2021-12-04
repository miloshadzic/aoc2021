package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"

	"github.com/miloshadzic/aoc2021/in"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("internal/inputs/day3.txt")
	check(err)
	defer f.Close()

	mcb := [12]int64{0, 0, 0, 0, 0}
	buf := make([]byte, 1)
	gamma, epsilon := 0, 0

	for i := 0; i < 12; i++ {
		f.Seek(int64(i), 0)

		for j := 0; ; j += 13 {
			_, err := f.ReadAt(buf, int64(i+j))
			if err != nil {
				if err == io.EOF {
					break
				}

				panic(err)
			}

			if buf[0] == '1' {
				mcb[i]++
			}
		}
	}

	s := bufio.NewScanner(f)
	total := 0

	for s.Scan() {
		total++
	}

	bits := [12]string{"0", "0", "0", "0", "0"}

	for i, b := range mcb {
		if b > int64(total/2) {
			gamma = gamma | (0b00001 << (11 - i))
			bits[i] = "1"
		}
	}

	epsilon = 0b111111111111 ^ gamma

	fmt.Println(gamma * epsilon)

	// Going to do this in a shittier way so I can do it on time

	lines := in.GetLines("day3")
	sort.Strings(lines)
	lines = lines[1:]
	selected := lines

	pos := 0

	for len(selected) > 1 {
		ones, rem := 0, len(selected)
		var filtered []string

		for i := range selected {
			if string(selected[i][pos]) == "1" {
				ones++
			}
		}

		zeroes := rem - ones

		if ones >= zeroes {
			for i := range selected {
				if string(selected[i][pos]) != "0" {
					filtered = append(filtered, selected[i])
				}
			}
		} else {
			for i := range selected {
				if string(selected[i][pos]) != "1" {
					filtered = append(filtered, selected[i])
				}
			}
		}

		selected = filtered
		pos++
	}

	oxy, err := strconv.ParseInt(selected[0], 2, 64)

	// FML

	selected = lines
	pos = 0

	for len(selected) > 1 {
		zeroes, rem := 0, len(selected)
		var filtered []string

		for i := range selected {
			if string(selected[i][pos]) == "0" {
				zeroes++
			}
		}

		ones := rem - zeroes

		if zeroes > ones {
			for i := range selected {
				if string(selected[i][pos]) != "0" {
					filtered = append(filtered, selected[i])
				}
			}
		} else {
			for i := range selected {
				if string(selected[i][pos]) != "1" {
					filtered = append(filtered, selected[i])
				}
			}
		}

		selected = filtered
		pos++
	}

	co2, err := strconv.ParseInt(selected[0], 2, 64)

	fmt.Println(oxy * co2)
}

func revBits(i int, bits [12]string) string {
	if bits[i] == "1" {
		return "0"
	} else {
		return "1"
	}
}
