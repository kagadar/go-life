package game

import (
	"strings"

	"github.com/kagadar/go-pipeline/slices"
)

type Point struct {
	x, y int
}

func clamp(b Board, p Point) Point {
	if p.x < 0 {
		p.x = b.width() + p.x
	} else {
		p.x = p.x % b.width()
	}
	if p.y < 0 {
		p.y = b.height() + p.y
	} else {
		p.y = p.y % b.height()
	}
	return p
}

func neighbours(b Board, p Point, fn func(Point)) {
	for x := p.x - 1; x <= p.x+1; x++ {
		for y := p.y - 1; y <= p.y+1; y++ {
			if x == p.x && y == p.y {
				continue
			}
			fn(clamp(b, Point{x, y}))
		}
	}
}

type State [][]bool

func (s State) String() string {
	return strings.Join(slices.Transform(s, func(row []bool) string {
		return strings.Join(slices.Transform(row, func(cell bool) string {
			if cell {
				return "■"
			} else {
				return "□"
			}
		}), " ")
	}), "\n")
}

type Board interface {
	width() int
	height() int

	State() State
	Tick()
}
