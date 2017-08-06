package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/pkg/term"
)

const (
	boxTop      = "\u2500"
	boxTopLeft  = "\u250c"
	boxTopRight = "\u2510"

	boxBottom      = "\u2500"
	boxBottomLeft  = "\u2514"
	boxBottomRight = "\u2518"

	boxLeft  = "\u2502"
	boxRight = "\u2502"

	block     = "\u2588"
	blockUp   = "\u2580"
	blockDown = "\u2584"
)

var stdin *term.Term

var cols, rows int

var (
	tooSmall    bool
	topPadding  int
	leftPadding int
)

func init() {
	fmt.Print("\x1b[?1049h") // tput smcup
	fmt.Print("\x1b[?25l")   // tput civis

	var err error
	stdin, err = term.Open("/dev/stdin")
	if err != nil {
		panic(err)
	}
	term.CBreakMode(stdin)
}

func restoreTerm() {
	stdin.Restore()
	fmt.Print("\x1b[H\x1b[2J") // clear
	fmt.Print("\x1b[?25h")     // tput cvvis
	fmt.Print("\x1b[?1049l")   // tput rmcup
}

func watchDimensions() {
	winch := make(chan os.Signal, 1)
	signal.Notify(winch, syscall.SIGWINCH)
	for {
		<-winch
		resize()
	}
}

func watchInput() {
	for {
		var input [3]byte
		os.Stdin.Read(input[:])

		switch strings.Trim(string(input[:]), "\x00") {
		case "\x1b[A", "w", "k":
			if snake.lastDirection != down {
				snake.direction = up
			}
		case "\x1b[B", "s", "j":
			if snake.lastDirection != up {
				snake.direction = down
			}
		case "\x1b[C", "d", "l":
			if snake.lastDirection != left {
				snake.direction = right
			}
		case "\x1b[D", "a", "h":
			if snake.lastDirection != right {
				snake.direction = left
			}
		}
	}
}

func resize() {
	// update terminal dimensions
	var err error
	cols, rows, err = terminal.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}

	// account for borders and newline
	tooSmall = rows < boardHeight/2+3 || cols < boardWidth+2

	topPadding = (rows - boardHeight/2 - 3) / 2
	leftPadding = (cols - boardWidth - 2) / 2

	// clear the screen
	fmt.Print("\x1b[H\x1b[2J")
	draw()
}
