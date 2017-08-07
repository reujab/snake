package main

import (
	"flag"
	"image"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var start = time.Now()

var tickInterval time.Duration

type state byte

var (
	boardWidth  int
	boardHeight int
)

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
	direction     direction
	lastDirection direction
	body          []image.Point
}

var snake = Snake{
	// initialize head
	body: []image.Point{{}},
}

// Food represents an apple.
type Food struct {
	pos image.Point
}

var food Food

func init() {
	flag.IntVar(&boardWidth, "width", 32, "the width of the board")
	flag.IntVar(&boardHeight, "height", 32, "the height of the board")
	flag.DurationVar(&tickInterval, "tick-interval", time.Millisecond*100, "the tick interval")
	flag.Parse()
	if boardHeight%2 != 0 {
		panic("board height must be even")
	}
}

func main() {
	go watchDimensions()
	go watchInput()

	resetFood()
	resize()
	go func() {
		for {
			tick()
			time.Sleep(tickInterval)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interrupt
	restoreTerm()
}
