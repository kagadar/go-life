package game

import (
	"runtime"
	"sync"
)

type concurrentBoard struct {
	w, h       int
	state      State
	weights    [][]byte
	work       chan func()
	weightPool sync.WaitGroup
	evolvePool sync.WaitGroup
}

func NewConcurrentBoard(state State) Board {
	board := &concurrentBoard{w: len(state[0]), h: len(state), state: state.Clone(), weights: make([][]byte, len(state)), work: make(chan func(), max(runtime.NumCPU()-1, 1))}
	for y := range board.h {
		board.weights[y] = make([]byte, board.w)
	}
	for range cap(board.work) {
		go func() {
			for {
				(<-board.work)()
			}
		}()
	}
	return board
}

func (c *concurrentBoard) Snapshot() State {
	return c.state.Clone()
}

func (c *concurrentBoard) String() string {
	return c.state.String()
}

func (c *concurrentBoard) Tick() {
	workers := cap(c.work)
	c.weightPool.Add(workers)
	c.evolvePool.Add(workers)
	for i := range workers {
		c.work <- func() {
			i, w, h, state, weights := i, c.w, c.h, c.state, c.weights
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
			c.weightPool.Done()
			c.weightPool.Wait()
			for y := i; y < h; y += workers {
				sRow := state[y]
				wRow := weights[y]
				for x, weight := range wRow {
					sRow[x] = weight == 3 || sRow[x] && weight == 2
					wRow[x] = 0
				}
			}
			c.evolvePool.Done()
		}
	}
	c.evolvePool.Wait()
}
