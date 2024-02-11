package game

import (
	"testing"
)

func TestSlice2_Concurrent_2D(t *testing.T) {
	testBoard(t, NewSlice2_2D_Concurrent)
}

func BenchmarkNewSlice2_Conncurrent_2D(b *testing.B) {
	benchmarkNew(b, NewSlice2_2D_Concurrent)
}

func BenchmarkSlice2_Concurrent_2D(b *testing.B) {
	benchmarkBoard(b, NewSlice2_2D_Concurrent)
}
