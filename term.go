package main

import (
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
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
			draw()
		}
	}()
}

func updateSize() {
	var err error
	cols, rows, err = terminal.GetSize(int(os.Stdin.Fd()))
	die(err)
}
