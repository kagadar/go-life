package game

import (
	"runtime"
	"sync"
)

type slice2_2D_concurrent struct {
	w, h       int
	state      State
	weights    [][]byte
	work       chan func()
	weightPool sync.WaitGroup
	evolvePool sync.WaitGroup
}

func NewSlice2_2D_Concurrent(s State) Board {
	b := slice2_2D_concurrent{w: len(s[0]), h: len(s), state: s.Clone(), weights: make([][]byte, len(s)), work: make(chan func(), max(runtime.NumCPU()-1, 1))}
	for y := range b.h {
		b.weights[y] = make([]byte, b.w)
	}
	for range cap(b.work) {
		go func() {
			for {
				(<-b.work)()
			}
		}()
	}
	return &b
}

func (b *slice2_2D_concurrent) Snapshot() State {
	return b.state.Clone()
}

func (b *slice2_2D_concurrent) String() string {
	return b.state.String()
}

func (b *slice2_2D_concurrent) Tick() {
	workers := cap(b.work)
	b.weightPool.Add(workers)
	b.evolvePool.Add(workers)
	for i := range workers {
		b.work <- func() {
			i, w, h, state, weights := i, b.w, b.h, b.state, b.weights
			for y := i; y < h; y += workers {
				row := weights[y]
				for x := range state[y] {
					neighbours(w, h, Point{x, y}, func(p Point) {
						if state[p.y][p.x] {
							row[x]++
						}
					})
				}
			}
			b.weightPool.Done()
			b.weightPool.Wait()
			for y := i; y < h; y += workers {
				sRow := state[y]
				wRow := weights[y]
				for x, weight := range wRow {
					sRow[x] = weight == 3 || sRow[x] && weight == 2
					wRow[x] = 0
				}
			}
			b.evolvePool.Done()
		}
	}
	b.evolvePool.Wait()
}
