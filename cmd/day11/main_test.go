package main

import (
	"testing"
)

func BenchmarkDay11(b *testing.B) {
	cave := InitCave("day11")

	for i := 0; i < b.N; i++ {
		cave.Next()
	}
}
