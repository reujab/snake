package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func resetFood() {
	food.pos.X = rand.Intn(boardWidth)
	food.pos.Y = rand.Intn(boardHeight)
	drawFood()
}

func tick() {
	if gameState != stateRunning || tooSmall {
		return
	}

	// erase old snake
	fmt.Printf("\x1b[%d;%dH ", topPadding+2+snake.pos.Y/2, leftPadding+2+snake.pos.X)

	// update snake position
	lastPos := snake.pos
	switch snake.direction {
	case right:
		snake.pos.X++
	case left:
		snake.pos.X--
	case down:
		snake.pos.Y++
	case up:
		snake.pos.Y--
	}

	// check if the player lost
	if snake.pos.X < 0 || snake.pos.X >= boardWidth || snake.pos.Y < 0 || snake.pos.Y >= boardHeight {
		gameState = stateOver
		snake.pos = lastPos
		drawSnake()
		drawGameOver()
	} else {
		drawSnake()
	}

	// check if snake is colliding with food
	if snake.pos.Eq(food.pos) {
		resetFood()
	}

	// check if snake and food occupy the same cell
	if snake.pos.X == food.pos.X && snake.pos.Y/2 == food.pos.Y/2 || lastPos.X == food.pos.X && lastPos.Y/2 == food.pos.Y/2 {
		drawFood()
	}
}
