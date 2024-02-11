package game

type slice1_Unroll struct {
	w, h, w2, wi, wh, hi, whi, whiu int
	state                           []byte
}

func NewSlice1_Unroll(s State) Board {
	b := slice1_Unroll{w: len(s[0]), h: len(s)}

	// precalculate useful numbers to navigate cells
	b.w2, b.wi, b.wh, b.hi = b.w+b.w, b.w-1, b.w*b.h, b.h-1
	b.whi = b.w * b.hi
	b.whiu = b.whi - b.w

	b.state = make([]byte, b.wh)
	for y := range b.h {
		for x := range b.w {
			if s[y][x] {
				b.state[y*b.w+x] = aliveMask
			}
		}
	}
	return &b
}

func (b *slice1_Unroll) Snapshot() State {
	out := make(State, b.h)
	for y := range b.h {
		out[y] = make([]bool, b.w)
		for x := range b.w {
			out[y][x] = b.state[y*b.w+x] > 0
		}
	}
	return out
}

func (b *slice1_Unroll) String() string {
	return b.Snapshot().String()
}

func (b *slice1_Unroll) Tick() {
	w, w2, wi, wh, hi, whi, whiu, state := b.w, b.w2, b.wi, b.wh, b.hi, b.whi, b.whiu, b.state
	/*
		00 01 02 03 04
		05 06 07 08 09
		10 11 12 13 14
		15 16 17 18 19
		20 21 22 23 24
	*/
	if state[0]&aliveMask > 0 {
		// 0
		state[1]++     // 1
		state[w-1]++   // 4
		state[w]++     // 5
		state[w+1]++   // 6
		state[w2-1]++  // 9
		state[whi]++   // 20
		state[whi+1]++ // 21
		state[wh-1]++  // 24
	}
	if state[whi]&aliveMask > 0 {
		// 20
		state[0]++      // 0
		state[1]++      // 1
		state[w-1]++    // 4
		state[whiu]++   // 15
		state[whiu+1]++ // 16
		state[whi-1]++  // 19
		state[whi+1]++  // 21
		state[wh-1]++   // 24
	}
	if state[w-1]&aliveMask > 0 {
		// 4
		state[0]++    // 0
		state[w-2]++  // 3
		state[w]++    // 5
		state[w2-2]++ // 8
		state[w2-1]++ // 9
		state[whi]++  // 20
		state[wh-2]++ // 23
		state[wh-1]++ // 24
	}
	if state[wh-1]&aliveMask > 0 {
		// 24
		state[0]++     // 0
		state[w-2]++   // 3
		state[w-1]++   // 4
		state[whiu]++  // 15
		state[whi-2]++ // 18
		state[whi-1]++ // 19
		state[whi]++   // 20
		state[wh-2]++  //23
	}
	for x := 1; x < wi; x++ {
		whix := whi + x
		if state[x]&aliveMask > 0 {
			wdx := w + x
			// 1
			state[x-1]++    // 0
			state[x+1]++    // 2
			state[wdx-1]++  // 5
			state[wdx]++    // 6
			state[wdx+1]++  // 7
			state[whix-1]++ // 20
			state[whix]++   // 21
			state[whix+1]++ // 22
		}
		if state[whix]&aliveMask > 0 {
			whiux := whiu + x
			// 21
			state[x-1]++     // 0
			state[x]++       // 1
			state[x+1]++     // 2
			state[whiux-1]++ // 15
			state[whiux]++   // 16
			state[whiux+1]++ // 17
			state[whix-1]++  // 20
			state[whix+1]++  // 22
		}
	}
	for y := 1; y < hi; y++ {
		wy := w * y
		wyu, wyd := wy-w, wy+w
		if state[wy]&aliveMask > 0 {
			// 5
			state[wyu]++     // 0
			state[wyu+1]++   // 1
			state[wy-1]++    // 4
			state[wy+1]++    // 6
			state[wyd-1]++   // 9
			state[wyd]++     // 10
			state[wyd+1]++   // 11
			state[wyd+w-1]++ // 14
		}
		for x := 1; x < wi; x++ {
			wyx, wyux, wydx := wy+x, wyu+x, wyd+x
			if state[wyx]&aliveMask > 0 {
				// 6
				state[wyux-1]++ // 0
				state[wyux]++   // 1
				state[wyux+1]++ // 2
				state[wyx-1]++  // 5
				state[wyx+1]++  // 7
				state[wydx-1]++ // 10
				state[wydx]++   // 11
				state[wydx+1]++ // 12
			}
		}
		if state[wyd-1]&aliveMask > 0 {
			wydd := wyd + w
			// 9
			state[wyu]++    // 0
			state[wy-2]++   // 3
			state[wy-1]++   // 4
			state[wy]++     // 5
			state[wyd-2]++  // 8
			state[wyd]++    // 10
			state[wydd-2]++ // 13
			state[wydd-1]++ // 14
		}
	}
	for i, cell := range state {
		alive := cell&aliveMask > 0
		weight := cell & weightMask
		if weight == 3 || alive && weight == 2 {
			state[i] = aliveMask
		} else {
			state[i] = 0
		}
	}
}
