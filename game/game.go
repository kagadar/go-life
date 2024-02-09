package game

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"strings"

	pslices "github.com/kagadar/go-pipeline/slices"
)

type Point struct {
	x, y int
}

func clamp(w, h int, p Point) Point {
	if p.x < 0 {
		p.x = w + p.x
	} else {
		p.x = p.x % w
	}
	if p.y < 0 {
		p.y = h + p.y
	} else {
		p.y = p.y % h
	}
	return p
}

func neighbours(w, h int, p Point, fn func(Point)) {
	for x := p.x - 1; x <= p.x+1; x++ {
		for y := p.y - 1; y <= p.y+1; y++ {
			if x == p.x && y == p.y {
				continue
			}
			fn(clamp(w, h, Point{x, y}))
		}
	}
}

type State [][]bool

func (s State) String() string {
	return strings.Join(pslices.Transform(s, func(row []bool) string {
		return strings.Join(pslices.Transform(row, func(cell bool) string {
			if cell {
				return "■"
			} else {
				return "□"
			}
		}), " ")
	}), "\n")
}

func (s State) Clone() State {
	return pslices.Transform(s, func(row []bool) []bool {
		return slices.Clone(row)
	})
}

func RandomState(width, height int, chance float64) State {
	state := make(State, height)
	for y := range height {
		state[y] = make([]bool, width)
	}
	for y := range height {
		for x := range width {
			state[y][x] = rand.Float64() <= chance
		}
	}
	return state
}

type Board interface {
	fmt.Stringer
	Snapshot() State
	Tick()
}
