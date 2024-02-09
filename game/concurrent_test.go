package game

import (
	"testing"
)

func TestConcurrentBoard(t *testing.T) {
	testBoard(t, NewConcurrentBoard)
}

func BenchmarkNewConncurrentBoard(b *testing.B) {
	benchmarkNew(b, NewConcurrentBoard)
}

func BenchmarkConcurrentBoard(b *testing.B) {
	benchmarkBoard(b, NewConcurrentBoard)
}
