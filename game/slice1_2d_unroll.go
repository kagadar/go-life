package game

import "github.com/kagadar/go-pipeline/slices"

type slice1_2D_unroll struct {
	w, h, wi, wii, hi int
	state             [][]byte
}

func NewSlice1_2D_Unroll(s State) Board {
	b := slice1_2D_unroll{w: len(s[0]), h: len(s), state: make([][]byte, len(s))}
	b.wi, b.wii, b.hi = b.w-1, b.w-2, b.h-1
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

func (b *slice1_2D_unroll) Snapshot() State {
	return slices.Transform(b.state, func(row []byte) []bool {
		return slices.Transform(row, func(cell byte) bool {
			return cell&aliveMask > 0
		})
	})
}

func (b *slice1_2D_unroll) String() string {
	return b.Snapshot().String()
}

func (b *slice1_2D_unroll) Tick() {
	wi, wii, hi, state := b.wi, b.wii, b.hi, b.state
	row0, row1, rowhi, rowhiu := state[0], state[1], state[hi], state[hi-1]
	/*
		(0,0) (0,1) (0,2) (0,3) (0,4)
		(1,0) (1,1) (1,2) (1,3) (1,4)
		(2,0) (2,1) (2,2) (2,3) (2,4)
		(3,0) (3,1) (3,2) (3,3) (3,4)
		(4,0) (4,1) (4,2) (4,3) (4,4)
	*/
	if row0[0]&aliveMask > 0 {
		// (0,0)
		row0[1]++   // (0,1)
		row0[wi]++  // (0,4)
		row1[0]++   // (1,0)
		row1[1]++   // (1,1)
		row1[wi]++  // (1,4)
		rowhi[0]++  // (4,0)
		rowhi[1]++  // (4,1)
		rowhi[wi]++ // (4,4)
	}
	if row0[wi]&aliveMask > 0 {
		// (0,4)
		row0[wii]++  // (0,3)
		row0[0]++    // (0,0)
		row1[wii]++  // (1,3)
		row1[wi]++   // (1,4)
		row1[0]++    // (1,0)
		rowhi[wii]++ // (4,3)
		rowhi[wi]++  // (4,4)
		rowhi[0]++   // (4,0)
	}
	if rowhi[0]&aliveMask > 0 {
		// (4,0)
		rowhi[1]++   // (4,1)
		rowhi[wi]++  // (4,4)
		rowhiu[0]++  // (3,0)
		rowhiu[1]++  // (3,1)
		rowhiu[wi]++ // (3,4)
		row0[0]++    // (0,0)
		row0[1]++    // (0,1)
		row0[wi]++   // (0,4)
	}
	if rowhi[wi]&aliveMask > 0 {
		// (4,4)
		rowhi[wii]++  // (4,3)
		rowhi[0]++    // (4,0)
		rowhiu[wii]++ // (3,3)
		rowhiu[wi]++  // (3,4)
		rowhiu[0]++   // (3,0)
		row0[wii]++   // (0,3)
		row0[wi]++    // (0,4)
		row0[0]++     // (0,0)
	}
	for x := 1; x < wi; x++ {
		before, after := x-1, x+1
		if row0[x]&aliveMask > 0 {
			// (0,1)
			row0[before]++  // (0,0)
			row0[after]++   // (0,2)
			row1[before]++  // (1,0)
			row1[x]++       // (1,1)
			row1[after]++   // (1,2)
			rowhi[before]++ // (4,0)
			rowhi[x]++      // (4,1)
			rowhi[after]++  // (4,2)
		}
		if rowhi[x]&aliveMask > 0 {
			// (4,1)
			rowhi[before]++  // (4,0)
			rowhi[after]++   // (4,2)
			rowhiu[before]++ // (3,0)
			rowhiu[x]++      // (3,1)
			rowhiu[after]++  // (3,2)
			row0[before]++   // (0,0)
			row0[x]++        // (0,1)
			row0[after]++    // (0,2)
		}
	}
	for y := 1; y < hi; y++ {
		rowu, row, rowd := state[y-1], state[y], state[y+1]
		if row[0]&aliveMask > 0 {
			// (1,0)
			row[1]++   // (1,1)
			row[wi]++  // (1,4)
			rowu[0]++  // (0,0)
			rowu[1]++  // (0,1)
			rowu[wi]++ // (0,4)
			rowd[0]++  // (2,0)
			rowd[1]++  // (2,1)
			rowd[wi]++ // (2,4)
		}
		for x := 1; x < wi; x++ {
			if row[x]&aliveMask > 0 {
				before, after := x-1, x+1
				// (1,1)
				row[before]++  // (1,0)
				row[after]++   // (1,2)
				rowu[before]++ // (0,0)
				rowu[x]++      // (0,1)
				rowu[after]++  // (0,2)
				rowd[before]++ // (2,0)
				rowd[x]++      // (2,1)
				rowd[after]++  // (2,2)
			}
		}
		if row[wi]&aliveMask > 0 {
			// (1,4)
			row[wii]++  // (1,3)
			row[0]++    // (1,0)
			rowu[wii]++ // (0,3)
			rowu[wi]++  // (0,4)
			rowu[0]++   // (0,0)
			rowd[wii]++ // (2,3)
			rowd[wi]++  // (2,4)
			rowd[0]++   // (2,0)
		}
	}
	for _, row := range state {
		for x, cell := range row {
			if cell < 3 {
				row[x] = 0
				continue
			}
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
