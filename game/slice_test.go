package game

import (
	"testing"
)

func TestSliceBoard(t *testing.T) {
	testBoard(t, NewSliceBoard)
}

func BenchmarkNewSliceBoard(b *testing.B) {
	benchmarkNew(b, NewSliceBoard)
}

func BenchmarkSliceBoard(b *testing.B) {
	benchmarkBoard(b, NewSliceBoard)
}
