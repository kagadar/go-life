package game

import (
	"github.com/kagadar/go-set"
)

type mapBoard struct {
	w, h    int
	cells   set.Set[Point]
	weights map[Point]byte
}

func NewMapBoard(state State) Board {
	board := &mapBoard{w: len(state[0]), h: len(state), cells: set.Set[Point]{}, weights: map[Point]byte{}}
	for y, row := range state {
		for x, cell := range row {
			if cell {
				board.cells.Put(Point{x, y})
			}
		}
	}
	return board
}

func (m *mapBoard) Snapshot() State {
	out := make(State, m.h)
	for y := range m.h {
		out[y] = make([]bool, m.w)
	}
	for p := range m.cells {
		out[p.y][p.x] = true
	}
	return out
}

func (m *mapBoard) String() string {
	return m.Snapshot().String()
}

func (m *mapBoard) Tick() {
	for p := range m.cells {
		//lint:ignore SA4018 a live cell with no neighbours should still be present in the weights.
		m.weights[p] = m.weights[p]
		neighbours(m.w, m.h, p, func(adj Point) {
			m.weights[adj]++
		})
	}
	for p, weight := range m.weights {
		if m.cells.Has(p) {
			if weight < 2 || weight > 3 {
				delete(m.cells, p)
			}
		} else {
			if weight == 3 {
				m.cells.Put(p)
			}
		}
		delete(m.weights, p)
	}
}
