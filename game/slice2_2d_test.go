package game

import (
	"testing"
)

func TestSlice2_2D(t *testing.T) {
	testBoard(t, NewSlice2_2D)
}

func BenchmarkNewSlice2_2D(b *testing.B) {
	benchmarkNew(b, NewSlice2_2D)
}

func BenchmarkSlice2_2D(b *testing.B) {
	benchmarkBoard(b, NewSlice2_2D)
}
