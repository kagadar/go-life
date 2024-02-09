package game

import (
	"strings"

	"github.com/kagadar/go-pipeline/slices"
)

const (
	aliveMask  byte = 0b10000000
	weightMask byte = 0b01111111
)

type bitBoard struct {
	w, h  int
	state [][]byte
}

func NewBitBoard(state State) Board {
	board := &bitBoard{w: len(state[0]), h: len(state), state: make([][]byte, len(state))}
	for y := range board.h {
		board.state[y] = make([]byte, board.w)
		for x := range board.w {
			if state[y][x] {
				board.state[y][x] = aliveMask
			}
		}
	}
	return board
}

func (b *bitBoard) Snapshot() State {
	return slices.Transform(b.state, func(row []byte) []bool {
		return slices.Transform(row, func(cell byte) bool {
			return cell&aliveMask > 0
		})
	})
}

func (b *bitBoard) String() string {
	return strings.Join(slices.Transform(b.state, func(row []byte) string {
		return strings.Join(slices.Transform(row, func(cell byte) string {
			if cell > 0 {
				return "■"
			} else {
				return "□"
			}
		}), " ")
	}), "\n")
}

func (b *bitBoard) Tick() {
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
