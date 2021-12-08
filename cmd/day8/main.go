package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func diff(a, b string) []rune {
	var out []rune
	m := make(map[rune]bool)

	for _, item := range b {
		m[item] = true
	}

	for _, item := range a {
		if _, ok := m[item]; !ok {
			out = append(out, item)
		}
	}

	return out
}

func rDiff(a, b []rune) []rune {
	var out []rune
	m := make(map[rune]bool)

	for _, item := range b {
		m[item] = true
	}

	for _, item := range a {
		if _, ok := m[item]; !ok {
			out = append(out, item)
		}
	}

	return out
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

// 1 = len(2)
// 7 = len(3)
// 8 = len(7)
// 4 = len(4)
// a = 7 - 1
// bd = 4 - 7
// e = len(6) - 4 - a // where len(1)
// g = 8 - 4 - e - a
// 9 = 4 + a + g

func main() {
	sum := 0
	f, _ := os.Open("internal/inputs/day8.txt")

	s := bufio.NewScanner(f)

	for s.Scan() {
		inOut := strings.Split(s.Text(), "| ")
		all := inOut[0] + inOut[1]
		// I am using memory luxuriously this time
		digits := make(map[string]int)
		revDig := make(map[int]string)

		strs := make(map[string]bool)
		temp := make(map[int][]string)

		pos := make(map[rune]rune)

		for _, d := range strings.Split(all, " ") {
			if len(d) == 0 {
				continue
			}
			sorted := SortString(strings.Replace(d, " ", "", -1))
			if _, ok := strs[sorted]; !ok {
				strs[sorted] = true
				temp[len(sorted)] = append(temp[len(sorted)], sorted)

				switch len(sorted) {
				case 2:
					digits[sorted] = 1
					revDig[1] = sorted
				case 3:
					digits[sorted] = 7
					revDig[7] = sorted
				case 4:
					digits[sorted] = 4
					revDig[4] = sorted
				case 7:
					digits[sorted] = 8
					revDig[8] = sorted
				}
			}
		}

		// We "need" to find 2 first. It's going to be the one for which
		// X - 7 - 4 is of length 2.
		var eg string
		for _, s := range temp[5] { // 2 is somewhere here
			// LOLCODE
			t := string(rDiff(diff(s, revDig[7]), []rune(revDig[4])))
			if len(t) == 2 {
				digits[s] = 2
				revDig[2] = s
				eg = t
			} else if len(t) == 1 {
				pos['g'] = []rune(t)[0]
			} else {
				panic("HALP")
			}
		}

		// To find b, we do 8 - 2 - 1
		pos['b'] = rDiff(diff(revDig[8], revDig[2]), []rune(revDig[1]))[0]

		// To find d, we do 4 - b - 1
		pos['d'] = rDiff(diff(revDig[4], revDig[1]), []rune{pos['b']})[0]

		pos['f'] = diff(revDig[1], revDig[2])[0]
		pos['c'] = diff(revDig[1], string(pos['f']))[0]
		pos['e'] = diff(eg, string(pos['g']))[0]

		revDig[0] = string(diff(revDig[8], string(pos['d'])))
		digits[revDig[0]] = 0

		revDig[6] = string(diff(revDig[8], string(pos['c'])))
		digits[revDig[6]] = 6

		revDig[9] = string(diff(revDig[8], string(pos['e'])))
		digits[revDig[9]] = 9

		revDig[3] = string(diff(revDig[8], string([]rune{pos['b'], pos['e']})))
		digits[revDig[3]] = 3

		revDig[5] = string(diff(revDig[6], string(pos['e'])))
		digits[revDig[5]] = 5

		var output string
		for _, w := range strings.Split(inOut[1], " ") {
			output += fmt.Sprint(digits[SortString(w)])
		}

		number, _ := strconv.Atoi(output)
		sum += number
	}

	fmt.Println(sum)
}
