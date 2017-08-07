package main

import (
	"fmt"
	"image"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func resetFood() {
	colliding := true
	for colliding {
		food.pos.X = rand.Intn(boardWidth)
		food.pos.Y = rand.Intn(boardHeight)

		// check if new food position collides with snake
		colliding = false
		for _, pos := range snake.body {
			if food.pos.Eq(pos) {
				colliding = true
				break
			}
		}
	}

	drawFood()
}

func tick() {
	if gameState != stateRunning || tooSmall {
		return
	}

	lastBody := make([]image.Point, len(snake.body))
	copy(lastBody, snake.body)

	for i := len(snake.body) - 1; i > 0; i-- {
		// erase old point
		fmt.Printf("\x1b[%d;%dH ", topPadding+2+snake.body[i].Y/2, leftPadding+2+snake.body[i].X)

		// update position
		snake.body[i] = snake.body[i-1]
	}

	// erase old snake head
	fmt.Printf("\x1b[%d;%dH ", topPadding+2+snake.body[0].Y/2, leftPadding+2+snake.body[0].X)

	// update snake head position
	switch snake.direction {
	case right:
		snake.body[0].X++
	case left:
		snake.body[0].X--
	case down:
		snake.body[0].Y++
	case up:
		snake.body[0].Y--
	}

	// check if the player lost
	var lost bool
	// check if the snake is colliding with itself
	for _, pos := range snake.body[1:] {
		if snake.body[0] == pos {
			lost = true
			break
		}
	}
	// check if snake is out of bounds
	lost = lost || snake.body[0].X < 0 || snake.body[0].X >= boardWidth || snake.body[0].Y < 0 || snake.body[0].Y >= boardHeight

	if lost {
		gameState = stateOver
		snake.body = lastBody
		drawSnake()
		drawGameOver()
	} else {
		drawSnake()
	}

	// check if snake is colliding with food
	if snake.body[0].Eq(food.pos) {
		snake.body = append(snake.body, lastBody[len(lastBody)-1])
		resetFood()
	}

	// check if snake and food occupy the same cell
	for _, pos := range snake.body {
		if pos.Y/2 == food.pos.Y/2 {
			drawFood()
			break
		}
	}

	snake.lastDirection = snake.direction
}
