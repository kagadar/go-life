package game

import (
	"testing"
)

func TestSlice1_Index(t *testing.T) {
	testBoard(t, NewSlice1_Index_Unroll)
}

func BenchmarkNewSlice1_Index(b *testing.B) {
	benchmarkNew(b, NewSlice1_Index_Unroll)
}

func BenchmarkSlice1_Index(b *testing.B) {
	benchmarkBoard(b, NewSlice1_Index_Unroll)
}
