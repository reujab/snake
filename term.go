package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
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

	blockUp   = "\u2580"
	blockDown = "\u2584"
)

const (
	boardWidth  = 70
	boardHeight = boardWidth / 2
)

var cols, rows int

var (
	tooSmall    bool
	topPadding  int
	leftPadding int
)

// detect terminal resizes
func init() {
	go func() {
		winch := make(chan os.Signal, 1)
		signal.Notify(winch, syscall.SIGWINCH)
		for {
			<-winch
			resize()
		}
	}()
}

func resize() {
	// update terminal dimensions
	var err error
	cols, rows, err = terminal.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}

	// account for borders and newline
	tooSmall = rows < boardHeight+3 || cols < boardWidth+2

	topPadding = (rows - boardHeight - 3) / 2
	leftPadding = (cols - boardWidth - 2) / 2

	// clear the screen
	fmt.Print("\x1b[H\x1b[2J")
	drawBoard()
}
