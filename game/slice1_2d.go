package game

import (
	"github.com/kagadar/go-pipeline/slices"
)

type slice1_2D struct {
	w, h  int
	state [][]byte
}

func NewSlice1_2D(s State) Board {
	b := slice1_2D{w: len(s[0]), h: len(s), state: make([][]byte, len(s))}
	for y := range b.h {
		b.state[y] = make([]byte, b.w)
		for x := range b.w {
			if s[y][x] {
				b.state[y][x] = aliveMask
			}
		}
	}
	return &b
}

func (b *slice1_2D) Snapshot() State {
	return slices.Transform(b.state, func(row []byte) []bool {
		return slices.Transform(row, func(cell byte) bool {
			return cell&aliveMask > 0
		})
	})
}

func (b *slice1_2D) String() string {
	return b.Snapshot().String()
}

func (b *slice1_2D) Tick() {
	w, h, state := b.w, b.h, b.state
	for y, row := range state {
		for x, cell := range row {
			if cell&aliveMask > 0 {
				neighbours(w, h, Point{x, y}, func(p Point) {
					state[p.y][p.x]++
				})
			}
		}
	}
	for _, row := range state {
		for x, cell := range row {
			alive := cell&aliveMask > 0
			weight := cell & weightMask
			if weight == 3 || alive && weight == 2 {
				row[x] = aliveMask
			} else {
				row[x] = 0
			}
		}
	}
}
