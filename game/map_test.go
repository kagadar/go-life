package game

import (
	"testing"
)

func TestMapBoard(t *testing.T) {
	testBoard(t, NewMapBoard)
}

func BenchmarkNewMapBoard(b *testing.B) {
	benchmarkNew(b, NewMapBoard)
}

func BenchmarkMapBoard(b *testing.B) {
	benchmarkBoard(b, NewMapBoard)
}
