package game

import (
	"testing"
)

func TestSliceBoard(t *testing.T) {
	HelperTestBoard(t, "sliceBoard", NewSliceBoard)
}

func BenchmarkNewSliceBoard(b *testing.B) {
	HelperBenchmarkNewBoard(b, NewSliceBoard)
}

func BenchmarkSliceBoardTick(b *testing.B) {
	HelperBenchmarkBoardTick(b, NewSliceBoard)
}
