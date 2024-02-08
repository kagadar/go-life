package game

type sliceBoard struct {
	w, h    int
	state   State
	weights [][]byte
}

func NewSliceBoard(state State) Board {
	board := &sliceBoard{w: len(state[0]), h: len(state), state: state, weights: make([][]byte, len(state))}
	for y := range board.h {
		board.weights[y] = make([]byte, board.w)
	}
	return board
}

func (s *sliceBoard) width() int {
	return s.w
}

func (s *sliceBoard) height() int {
	return s.h
}

func (s *sliceBoard) State() State {
	return s.state
}

func (s *sliceBoard) Tick() {
	for y, row := range s.state {
		for x, cell := range row {
			if cell {
				neighbours(s, Point{x, y}, func(p Point) {
					s.weights[p.y][p.x]++
				})
			}
		}
	}
	for y, row := range s.weights {
		for x, weight := range row {
			s.state[y][x] = s.state[y][x] && (weight == 2 || weight == 3) || !s.state[y][x] && weight == 3
			s.weights[y][x] = 0
		}
	}
}
