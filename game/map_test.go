package game

import (
	"testing"
)

func TestMapBoard(t *testing.T) {
	HelperTestBoard(t, "mapBoard", NewMapBoard)
}

func BenchmarkNewMapBoard(b *testing.B) {
	HelperBenchmarkNewBoard(b, NewMapBoard)
}

func BenchmarkMapBoardTick(b *testing.B) {
	HelperBenchmarkBoardTick(b, NewMapBoard)
}
