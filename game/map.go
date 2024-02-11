package game

import (
	"github.com/kagadar/go-set"
)

type mapBoard struct {
	w, h    int
	cells   set.Set[Point]
	weights map[Point]byte
}

func NewMapBoard(s State) Board {
	b := mapBoard{w: len(s[0]), h: len(s), cells: set.Set[Point]{}, weights: map[Point]byte{}}
	for y, row := range s {
		for x, cell := range row {
			if cell {
				b.cells.Put(Point{x, y})
			}
		}
	}
	return &b
}

func (b *mapBoard) Snapshot() State {
	out := make(State, b.h)
	for y := range b.h {
		out[y] = make([]bool, b.w)
	}
	for p := range b.cells {
		out[p.y][p.x] = true
	}
	return out
}

func (b *mapBoard) String() string {
	return b.Snapshot().String()
}

func (b *mapBoard) Tick() {
	for p := range b.cells {
		//lint:ignore SA4018 a live cell with no neighbours should still be present in the weights.
		b.weights[p] = b.weights[p]
		neighbours(b.w, b.h, p, func(adj Point) {
			b.weights[adj]++
		})
	}
	for p, weight := range b.weights {
		if b.cells.Has(p) {
			if weight < 2 || weight > 3 {
				delete(b.cells, p)
			}
		} else {
			if weight == 3 {
				b.cells.Put(p)
			}
		}
		delete(b.weights, p)
	}
}
