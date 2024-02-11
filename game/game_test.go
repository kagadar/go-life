package game

import (
	"fmt"
	"slices"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/kagadar/go-pipeline/maps"
	pslices "github.com/kagadar/go-pipeline/slices"
)

const (
	x = true
	o = false
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
		out = slices.Concat(out, pslices.Transform(in, func(in []bool) []bool {
			return slices.Clone(in)
		}))
	}
	return out
}

type test struct {
	name string
	in   State
	want State
}

type benchmark struct {
	name  string
	state State
}

var (
	tests = []test{
		{
			name: "lonely5x5",
			in: State{
				{o, o, o, o, o},
				{o, o, o, o, o},
				{o, o, x, o, o},
				{o, o, o, o, o},
				{o, o, o, o, o},
			},
			want: State{
				{o, o, o, o, o},
				{o, o, o, o, o},
				{o, o, o, o, o},
				{o, o, o, o, o},
				{o, o, o, o, o},
			},
		},
		{
			name: "glide5x5",
			in: State{
				{o, o, o, o, o},
				{o, o, x, o, o},
				{o, o, o, x, o},
				{o, x, x, x, o},
				{o, o, o, o, o},
			},
			want: State{
				{o, o, o, o, o},
				{o, o, o, o, o},
				{o, x, o, x, o},
				{o, o, x, x, o},
				{o, o, x, o, o},
			},
		},
	}
	benchmarkOut Board
	benchmarks   = slices.Concat(pslices.Flatten(pslices.Transform(tests, func(t test) []benchmark {
		return []benchmark{
			{
				name:  fmt.Sprintf("%s-in", t.name),
				state: t.in,
			},
			{
				name:  fmt.Sprintf("%s-want", t.name),
				state: t.want,
			},
		}
	})), []benchmark{
		{"false1000x1000", multiplyState(1000, 1000, State{{o}})},
		{"true1000x1000", multiplyState(1000, 1000, State{{x}})},
		{"highDensity999x999", multiplyState(333, 333, State{
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
	})
	boards = []struct {
		name       string
		newBoardFn newBoardFn
	}{
		{"map", NewMapBoard},
		{"slice1_unroll", NewSlice1_Unroll},
		{"slice1_2d_unroll", NewSlice1_2D_Unroll},
		{"slice1_2d", NewSlice1_2D},
		{"slice2_2d", NewSlice2_2D},
		{"slice2_2d_concurrent", NewSlice2_2D_Concurrent},
	}
)

type newBoardFn func(State) Board

func testBoard(t *testing.T, f newBoardFn) {
	t.Helper()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Helper()
			board := f(tc.in)
			board.Tick()
			if diff := cmp.Diff(board.Snapshot(), tc.want); diff != "" {
				t.Errorf("State() after Tick() unexpected diff (-got +want):\n%s", diff)
			}
		})
	}
}

func benchmarkNew(b *testing.B, f newBoardFn) {
	b.Helper()
	for _, tc := range benchmarks {
		b.Run(tc.name, func(b *testing.B) {
			b.Helper()
			var board Board
			b.ResetTimer()
			for range b.N {
				board = f(tc.state)
			}
			benchmarkOut = board
		})
	}
}

func benchmarkTick(b *testing.B, board Board) time.Duration {
	b.Helper()
	b.ResetTimer()
	for range b.N {
		board.Tick()
	}
	return b.Elapsed() / time.Duration(b.N)
}

func benchmarkBoard(b *testing.B, f newBoardFn) {
	b.Helper()
	for _, tc := range benchmarks {
		b.Run(tc.name, func(b *testing.B) {
			b.Helper()
			benchmarkTick(b, f(tc.state))
		})
	}
	for chance := 0.05; chance < 1; chance += 0.05 {
		b.Run(fmt.Sprintf("random %.f%%", chance*100), func(b *testing.B) {
			b.Helper()
			benchmarkTick(b, f(RandomState(1000, 1000, chance)))
		})
	}
}

const iterations = 25

func TestDrift(t *testing.T) {
	in := RandomState(25, 25, 0.4)
	out := map[string][]State{}
	for _, tc := range boards {
		b := tc.newBoardFn(in)
		for range iterations {
			b.Tick()
			out[tc.name] = append(out[tc.name], b.Snapshot())
		}
	}
	keys := maps.Keys(out)
	for i := range keys {
		for j := i + 1; j < len(keys); j++ {
			if cmp.Diff(out[keys[i]], out[keys[j]]) != "" {
				t.Errorf("Tick() produces different states between %s and %s", keys[i], keys[j])
			}
		}
	}
}

func BenchmarkAllBoards(b *testing.B) {
	type report struct {
		name    string
		results map[string]time.Duration
	}
	var rankings []report
	b.Run("fixed benchmarks", func(b *testing.B) {
		for _, tc := range benchmarks {
			r := report{name: tc.name, results: map[string]time.Duration{}}
			b.Run(r.name, func(b *testing.B) {
				for _, board := range boards {
					b.Run(board.name, func(b *testing.B) {
						r.results[board.name] = benchmarkTick(b, board.newBoardFn(tc.state))
					})
				}
			})
			rankings = append(rankings, r)
		}
	})
	b.Run("random benchmarks", func(b *testing.B) {
		for chance := 0.05; chance < 1; chance += 0.05 {
			r := report{name: fmt.Sprintf("random %.f%%", chance*100), results: map[string]time.Duration{}}
			state := RandomState(1000, 1000, chance)
			b.Run(r.name, func(b *testing.B) {
				for _, board := range boards {
					b.Run(board.name, func(b *testing.B) {
						r.results[board.name] = benchmarkTick(b, board.newBoardFn(state))
					})
				}
			})
			rankings = append(rankings, r)
		}
	})
	for _, report := range rankings {
		fmt.Printf("%q rankings:\n", report.name)
		maps.ValueSortedRange(report.results, func(name string, d time.Duration) error {
			fmt.Printf("\t%q: %s\n", name, d)
			return nil
		})
	}
}
