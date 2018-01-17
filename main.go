package main

import (
	"flag"
	"image"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	stateRunning state = iota
	stateOver

	right direction = iota
	left
	down
	up
)

var (
	start = time.Now()

	tickInterval time.Duration
	boardWidth   int
	boardHeight  int

	gameState = stateRunning

	snake = Snake{
		// initialize head
		body: []image.Point{{}},
	}

	food Food
)

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
