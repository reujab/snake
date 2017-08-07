package main

import "image"

type state byte
type direction byte

// Snake represents a snake.
type Snake struct {
	direction     direction
	lastDirection direction
	body          []image.Point
}

// Food represents an apple.
type Food struct {
	pos image.Point
}
