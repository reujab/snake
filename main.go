package main

import (
	"fmt"
	"image"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type state byte

const (
	stateRunning state = iota
	stateOver
)

var gameState = stateRunning

type direction byte

const (
	right direction = iota
	left
	down
	up
)

// Snake represents a snake.
type Snake struct {
	pos       image.Point
	direction direction
}

var snake Snake

func main() {
	// switch to alternate terminal screen
	fmt.Print("\x1b[?1049h") // tput smcup
	// hide the cursor
	fmt.Print("\x1b[?25l") // tput civis
	defer func() {
		fmt.Print("\x1b[?1049l") // tput rmcup
		fmt.Print("\x1b[?25h")   // tput cvvis
	}()

	drawBoard()
	go func() {
		for {
			tick()
			time.Sleep(time.Millisecond * 250)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interrupt
}
