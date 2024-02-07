package game

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

type Board interface {
	width() int
	height() int

	State() [][]bool
	Tick()
}
