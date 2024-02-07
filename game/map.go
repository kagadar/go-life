package game

import (
	"github.com/kagadar/go-set"
)

type mapBoard struct {
	w, h  int
	cells set.Set[Point]
}

func NewMapBoard(state [][]bool) Board {
	board := &mapBoard{w: len(state[0]), h: len(state), cells: set.Set[Point]{}}
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

func (m *mapBoard) State() [][]bool {
	state := make([][]bool, m.height())
	for i := 0; i < m.height(); i++ {
		state[i] = make([]bool, m.width())
	}
	for p := range m.cells {
		state[p.y][p.x] = true
	}
	return state
}

func (m *mapBoard) Tick() {

}
