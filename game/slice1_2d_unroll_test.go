package game

import (
	"testing"
)

func TestSlice1_2D_Unroll(t *testing.T) {
	testBoard(t, NewSlice1_2D_Unroll)
}

func BenchmarkNewSlice1_2D_Unroll(b *testing.B) {
	benchmarkNew(b, NewSlice1_2D_Unroll)
}

func BenchmarkSlice1_2D_Unroll(b *testing.B) {
	benchmarkBoard(b, NewSlice1_2D_Unroll)
}
