package game

type sliceBoard struct {
	w, h    int
	state   State
	weights [][]byte
}

func NewSliceBoard(state State) Board {
	board := &sliceBoard{w: len(state[0]), h: len(state), state: state.Clone(), weights: make([][]byte, len(state))}
	for y := range board.h {
		board.weights[y] = make([]byte, board.w)
	}
	return board
}

func (s *sliceBoard) Snapshot() State {
	return s.state.Clone()
}

func (s *sliceBoard) String() string {
	return s.state.String()
}

func (s *sliceBoard) Tick() {
	w, h, state, weights := s.w, s.h, s.state, s.weights
	for y, row := range state {
		for x, cell := range row {
			if cell {
				neighbours(w, h, Point{x, y}, func(p Point) {
					weights[p.y][p.x]++
				})
			}
		}
	}
	for y, wRow := range weights {
		sRow := state[y]
		for x, weight := range wRow {
			sRow[x] = weight == 3 || sRow[x] && weight == 2
		}
		clear(wRow)
	}
}
