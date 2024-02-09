package game

import (
	"testing"
)

func TestBitBoard(t *testing.T) {
	testBoard(t, NewBitBoard)
}

func BenchmarkNewBitBoard(b *testing.B) {
	benchmarkNew(b, NewBitBoard)
}

func BenchmarkBitBoard(b *testing.B) {
	benchmarkBoard(b, NewBitBoard)
}
