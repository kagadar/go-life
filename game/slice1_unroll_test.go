package game

import (
	"testing"
)

func TestSlice1_Unroll(t *testing.T) {
	testBoard(t, NewSlice1_Unroll)
}

func BenchmarkNewSlice1_Unroll(b *testing.B) {
	benchmarkNew(b, NewSlice1_Unroll)
}

func BenchmarkSlice1_Unroll(b *testing.B) {
	benchmarkBoard(b, NewSlice1_Unroll)
}
