package tutil

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

type Point struct {
	X, Y int
}

type Terminal struct {
	State    *term.State
	Position Point
}

func InitializeRaw() (Terminal, error) {
	state, err := term.MakeRaw(int(os.Stdin.Fd()))
	t := Terminal{}
	t.Position = Point{0, 0}
	t.State = state
	return t, err
}

func (t Terminal) CloseTerminal() {
	term.Restore(int(os.Stdin.Fd()), t.State)
}

func (Terminal) ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func (t Terminal) MoveCursorToHome() {
	fmt.Print("\033[H")
	t.Position.X = 0
	t.Position.Y = 0
}

func (t Terminal) MoveCursorTo(line int, column int) {
	fmt.Printf("\033[%d;%dH", line, column)
	t.Position.X = column
	t.Position.Y = line
}

func (t Terminal) CursorUpN(n int) {
	fmt.Printf("\033[%dA", n)
	t.Position.Y -= n
}

func (t Terminal) CursorDownN(n int) {
	fmt.Printf("\033[%dB", n)
	t.Position.Y += n
}

func (t Terminal) CursorRightN(n int) {
	fmt.Printf("\033[%dC", n)
	t.Position.X += n
}

func (t Terminal) CursorLeftN(n int) {
	fmt.Printf("\033[%dD", n)
	t.Position.X -= n
}

func (t Terminal) CursorDownBeginningN(n int) {
	fmt.Printf("\033[%dE", n)
	t.Position.X = 0
	t.Position.Y += n
}

func (t Terminal) CursorUpBeginningN(n int) {
	fmt.Printf("\033[%dF", n)
	t.Position.X = 0
	t.Position.Y -= n
}

func (Terminal) SaveCursorPos() {
	fmt.Print("\033[s")
}

func (Terminal) RestorCursorPos() {
	fmt.Print("\033[u")
}
