package game

import (
	"github.com/kagadar/go-set"
)

type mapBoard struct {
	w, h    int
	state   State
	cells   set.Set[Point]
	weights map[Point]byte
}

func NewMapBoard(state State) Board {
	board := &mapBoard{w: len(state[0]), h: len(state), state: state, cells: set.Set[Point]{}, weights: map[Point]byte{}}
	for y, row := range state {
		for x, cell := range row {
			if cell {
				board.cells.Put(Point{x, y})
			}
		}
	}
	return board
}

func (m *mapBoard) width() int {
	return m.w
}

func (m *mapBoard) height() int {
	return m.h
}

func (m *mapBoard) State() State {
	return m.state
}

func (m *mapBoard) Tick() {
	for p := range m.cells {
		neighbours(m, p, func(adj Point) {
			m.weights[adj]++
		})
	}
	for p, weight := range m.weights {
		if m.cells.Has(p) {
			if weight < 2 || weight > 3 {
				delete(m.cells, p)
				m.state[p.y][p.x] = false
			}
		} else {
			if weight == 3 {
				m.cells.Put(p)
				m.state[p.y][p.x] = true
			}
		}
		m.weights[p] = 0
	}
}
