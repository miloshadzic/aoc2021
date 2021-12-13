package main

import (
	"bufio"
	"os"
)

type State int

const (
	UNDISCOVERED State = iota
	DISCOVERED
	PROCESSED
)

type CaveSize int

const (
	SMALL_CAVE CaveSize = iota
	LARGE_CAVE
)

type Node struct {
	Name string
	Size CaveSize
}

type EdgeNode struct {
	ID    uint
	Name  string
	State State
	Next  *EdgeNode
}

type Graph struct {
	Nodes     map[string]int
	Edges     []*EdgeNode
	NVertices uint
	NEdges    uint
}

func main() {
	f, _ := os.Open("internal/inputs/day12_test.txt")
	s := bufio.NewScanner(f)

	for s.Scan() {
	}
}
