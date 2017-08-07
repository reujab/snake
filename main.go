package main

import (
	"image"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var start = time.Now()

type state byte

const (
	boardWidth  = 40
	boardHeight = boardWidth
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

func main() {
	go watchDimensions()
	go watchInput()

	resetFood()
	resize()
	go func() {
		for {
			tick()
			time.Sleep(time.Millisecond * 100)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interrupt
	restoreTerm()
}
