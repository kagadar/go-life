package game

import "github.com/kagadar/go-set"

type mapBoard struct {
	w, h            int
	living, updated set.Set[Point]
	state           [][]byte
}

func NewMapBoard(s State) Board {
	b := mapBoard{w: len(s[0]), h: len(s)}
	wh := b.w * b.h
	b.living = make(set.Set[Point], wh)
	b.updated = make(set.Set[Point], wh)
	b.state = make([][]byte, b.h)
	for y, row := range s {
		b.state[y] = make([]byte, b.w)
		for x, cell := range row {
			if cell {
				b.state[y][x] = aliveMask
				b.living.Put(Point{x, y})
			}
		}
	}
	return &b
}

func (b *mapBoard) Snapshot() State {
	out := make(State, b.h)
	for y := range b.h {
		out[y] = make([]bool, b.w)
	}
	for p := range b.living {
		out[p.y][p.x] = true
	}
	return out
}

func (b *mapBoard) String() string {
	return b.Snapshot().String()
}

func (b *mapBoard) Tick() {
	w, h, living, updated, state := b.w, b.h, b.living, b.updated, b.state
	for p := range living {
		updated.Put(p)
		neighbours(w, h, p, func(adj Point) {
			state[adj.y][adj.x]++
			updated.Put(adj)
		})
	}
	clear(living)
	for p := range updated {
		row := state[p.y]
		cell := row[p.x]
		if cell < 3 {
			row[p.x] = 0
			continue
		}
		alive := cell&aliveMask > 0
		weight := cell & weightMask
		if weight == 3 || alive && weight == 2 {
			row[p.x] = aliveMask
			living.Put(p)
		} else {
			row[p.x] = 0
		}
	}
	clear(b.updated)
}
