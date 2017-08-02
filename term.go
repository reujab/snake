package main

import (
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

func init() {
	updateSize()
}

// detect terminal resizes
func init() {
	go func() {
		winch := make(chan os.Signal, 1)
		signal.Notify(winch, syscall.SIGWINCH)
		for {
			<-winch
			updateSize()
			drawBoard()
		}
	}()
}

func updateSize() {
	var err error
	cols, rows, err = terminal.GetSize(int(os.Stdin.Fd()))
	die(err)
}
