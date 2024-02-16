package game

type slice1_index_unroll struct {
	w, h                       int
	lastCol, lastRow, lastCell int
	living                     []int
	state                      []byte
}

func NewSlice1_Index_Unroll(s State) Board {
	b := slice1_index_unroll{w: len(s[0]), h: len(s)}
	wh := b.w * b.h

	b.lastCol = b.w - 1
	b.lastRow = b.w * (b.h - 1)
	b.lastCell = wh - 1

	b.living = make([]int, 0, wh)
	b.state = make([]byte, wh)
	for y := range b.h {
		for x := range b.w {
			if s[y][x] {
				cell := y*b.w + x
				b.living = append(b.living, cell)
				b.state[cell] = aliveMask
			}
		}
	}
	return &b
}

func (b *slice1_index_unroll) Snapshot() State {
	out := make(State, b.h)
	for y := range b.h {
		out[y] = make([]bool, b.w)
	}
	for _, i := range b.living {
		out[i/b.w][i%b.w] = true
	}
	return out
}

func (b *slice1_index_unroll) String() string {
	return b.Snapshot().String()
}

func (b *slice1_index_unroll) Tick() {
	w, lastCol, lastRow, lastCell, living, state := b.w, b.lastCol, b.lastRow, b.lastCell, b.living, b.state
	/*
		00 01 02 03 04
		05 06 07 08 09
		10 11 12 13 14
		15 16 17 18 19
		20 21 22 23 24
	*/
	for _, i := range b.living {
		var adjs []int
		switch {
		case i == 0:
			// tl corner (0)
		case i == lastCol:
			// tr corner (4)
		case i == lastRow:
			// bl corner (20)
		case i == lastCell:
			// br corner (24)
		case i < w:
			// first row
		case i >= lastRow:
			// last row
		case i%w == 0:
			// first column
		case i%w == lastCol:
			// last column
		default:
			wu, wd := i-w, i+w
			adjs = []int{
				wu - 1, wu, wu + 1,
				i - 1, i + 1,
				wd - 1, wd, wd + 1,
			}
		}
		for _, adj := range adjs {
			state[adj]++
		}
	}
	living = living[0:0]
	for i, cell := range state {
		if cell < 3 {
			state[i] = 0
			continue
		}
		alive := cell&aliveMask > 0
		weight := cell & weightMask
		if weight == 3 || alive && weight == 2 {
			living = append(living, i)
			state[i] = aliveMask
		} else {
			state[i] = 0
		}
	}
	b.living = living
}
