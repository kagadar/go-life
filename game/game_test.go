package game

import (
	"slices"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const (
	x = true
	o = false
)

var (
	glide5x5frame1 = State{
		{o, o, o, o, o},
		{o, o, x, o, o},
		{o, o, o, x, o},
		{o, x, x, x, o},
		{o, o, o, o, o},
	}
	glide5x5frame2 = State{
		{o, o, o, o, o},
		{o, o, o, o, o},
		{o, x, o, x, o},
		{o, o, x, x, o},
		{o, o, x, o, o},
	}
	benchmarkOut Board
	benchmarks   = []struct {
		name  string
		state State
	}{
		{"glide5x5frame1", glide5x5frame1},
		{"glide5x5frame2", glide5x5frame2},
		{"false1000x1000", multiplyState(1000, 1000, State{{o}})},
		{"true1000x1000", multiplyState(1000, 1000, State{{x}})},
		{"highDensity3000x3000", multiplyState(1000, 1000, State{
			{x, x, o},
			{x, x, o},
			{o, o, o},
		})},
		{"lowDensity1000x1000", mutateState(multiplyState(1000, 1000, State{{o}}),
			Point{1, 0},
			Point{2, 1},
			Point{0, 2},
			Point{1, 2},
			Point{2, 2},
		)},
	}
)

func mutateState(state State, points ...Point) State {
	for _, p := range points {
		state[p.y][p.x] = true
	}
	return state
}

func multiplyState(width, height int, in State) (out State) {
	for y, row := range in {
		for range width {
			in[y] = slices.Concat(in[y], row)
		}
	}
	for range height {
		out = slices.Concat(out, slices.Clone(in))
	}
	return out
}

type newBoardFn func(State) Board

func HelperTestBoard(t *testing.T, name string, f newBoardFn) {
	t.Helper()
	board := f(glide5x5frame1)
	board.Tick()
	if diff := cmp.Diff(board.State(), glide5x5frame2); diff != "" {
		t.Errorf("%s.State() after Tick() unexpected diff (-got +want):\n%s", name, diff)
	}
}

func HelperBenchmarkNewBoard(b *testing.B, f newBoardFn) {
	b.Helper()
	for _, tc := range benchmarks {
		b.Run(tc.name, func(b *testing.B) {
			var board Board
			for range b.N {
				board = f(tc.state)
			}
			benchmarkOut = board
		})
	}
}

func HelperBenchmarkBoardTick(b *testing.B, f newBoardFn) {
	b.Helper()
	for _, tc := range benchmarks {
		b.Run(tc.name, func(b *testing.B) {
			board := f(tc.state)
			b.ResetTimer()
			for range b.N {
				board.Tick()
			}
		})
	}
}
