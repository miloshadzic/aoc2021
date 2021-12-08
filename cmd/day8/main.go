package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("internal/inputs/day8_test.txt")
	counts := make(map[int]int)

	s := bufio.NewScanner(f)

	for s.Scan() {
		inOut := strings.Split(s.Text(), " | ")
		out := inOut[1]

		fmt.Println(out)

		for _, d := range strings.Split(out, " ") {
			fmt.Println(d)
			switch len(d) {
			case 2:
				counts[1]++
			case 3:
				counts[7]++
			case 4:
				counts[4]++
			case 7:
				counts[8]++
			}
		}
	}

	sum := 0
	for _, v := range counts {
		sum += v
	}

	fmt.Println(sum)
}
