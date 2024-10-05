package tutil

import "fmt"

type Rect struct {
	X, Y, W, H int
}

func (t Terminal) DrawRect(r Rect) {
	t.MoveCursorTo(r.Y, r.X)
	fmt.Printf("%s", RoundTopLeft)
	for range r.W - 1 {
		fmt.Printf("%s", Horizontal)
	}
	fmt.Printf("%s", RoundTopRight)

	t.MoveCursorTo(r.Y+1, r.X)
	for i := range r.H - 2 {
		t.MoveCursorTo(r.Y+i+1, r.X)
		fmt.Printf("%s", Vertical)
		t.MoveCursorTo(r.Y+i+1, r.X+r.W)
		fmt.Printf("%s", Vertical)
	}

	t.MoveCursorTo(r.Y+r.H-1, r.X)
	fmt.Printf("%s", RoundBottomLeft)
	for range r.W - 1 {
		fmt.Printf("%s", Horizontal)
	}
	fmt.Printf("%s", RoundBottomRight)
}
