package main

import "fmt"

func tick() {
	if gameState != stateRunning {
		return
	}

	// erase old snake
	fmt.Printf("\x1b[%d;%dH ", topPadding+2+snake.pos.Y, leftPadding+2+snake.pos.X)

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
}
