package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

	mcb := [12]int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
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

	fmt.Println(total)
	fmt.Println(mcb)

	for i, b := range mcb {
		if b > int64(total/2) {
			gamma = gamma | (0b00001 << (11 - i))
		}
	}

	epsilon = 0b111111111111 ^ gamma

	fmt.Println(gamma)
	fmt.Println(epsilon)
	fmt.Println(gamma * epsilon)
}
