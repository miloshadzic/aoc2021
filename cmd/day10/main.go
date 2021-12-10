package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Stack struct {
	stack []rune
}

func (s *Stack) Push(d rune) {
	s.stack = append(s.stack, d)
}

func (s *Stack) Pop() (rune, error) {
	len := len(s.stack)
	if len == 0 {
		return -1, fmt.Errorf("Empty stack")
	}

	val := s.stack[len-1]
	s.stack = s.stack[:len-1]

	return val, nil
}

func (s *Stack) Peek() (rune, error) {
	len := len(s.stack)
	if len == 0 {
		return -1, fmt.Errorf("Empty stack")
	}

	val := s.stack[len-1]

	return val, nil
}

func open(d rune) bool {
	if d == '(' || d == '[' || d == '{' || d == '<' {
		return true
	}

	return false
}

func matching(d, o rune) bool {
	return o == closingDelimiter(d)
}

func closingDelimiter(d rune) rune {
	switch d {
	case '(':
		return ')'
	case '[':
		return ']'
	case '{':
		return '}'
	case '<':
		return '>'
	}

	return -1
}

var illegals = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var incompletes = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func main() {
	f, _ := os.Open("internal/inputs/day10.txt")
	s := bufio.NewScanner(f)
	var illegalScore int
	var incompleteScores []int

	for s.Scan() {
		stack := Stack{}
		illegal := false

		for _, d := range s.Text() {
			if open(d) {
				stack.Push(d)
			} else {
				top, _ := stack.Peek()

				if matching(top, d) {
					stack.Pop()
				} else {
					illegalScore += illegals[d]
					illegal = true
					break
				}
			}
		}

		if illegal {
			continue
		}

		score := 0
		for len(stack.stack) > 0 {
			val, _ := stack.Pop()
			score = score*5 + incompletes[closingDelimiter(val)]
		}
		incompleteScores = append(incompleteScores, score)
	}

	fmt.Println(illegalScore)

	sort.Ints(incompleteScores)
	fmt.Println(incompleteScores[len(incompleteScores)/2])
}
