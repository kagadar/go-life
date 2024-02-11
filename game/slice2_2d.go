package game

type slice2_2D struct {
	w, h    int
	state   State
	weights [][]byte
}

func NewSlice2_2D(s State) Board {
	b := slice2_2D{w: len(s[0]), h: len(s), state: s.Clone(), weights: make([][]byte, len(s))}
	for y := range b.h {
		b.weights[y] = make([]byte, b.w)
	}
	return &b
}

func (b *slice2_2D) Snapshot() State {
	return b.state.Clone()
}

func (b *slice2_2D) String() string {
	return b.state.String()
}

func (b *slice2_2D) Tick() {
	w, h, state, weights := b.w, b.h, b.state, b.weights
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
