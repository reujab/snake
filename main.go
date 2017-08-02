package main

import (
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
	go watchDimensions()
	go watchInput()

	resize()
	go func() {
		for {
			tick()
			time.Sleep(time.Millisecond * 250)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interrupt
	restoreTerm()
}
