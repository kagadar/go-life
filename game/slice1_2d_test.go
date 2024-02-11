package game

import (
	"testing"
)

func TestSlice1_2D(t *testing.T) {
	testBoard(t, NewSlice1_2D)
}

func BenchmarkNewSlice1_2D(b *testing.B) {
	benchmarkNew(b, NewSlice1_2D)
}

func BenchmarkSlice1_2D(b *testing.B) {
	benchmarkBoard(b, NewSlice1_2D)
}
