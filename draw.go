package main

import (
	"fmt"
	"strings"
)

func drawBoard() {
	if tooSmall {
		fmt.Print("terminal too small")
		return
	}

	drawBox()
	drawSnake()
	drawFood()
	if gameState == stateOver {
		drawGameOver()
	}
}

func drawBox() {
	fmt.Print(strings.Repeat("\n", topPadding))

	fmt.Print(strings.Repeat(" ", leftPadding))
	fmt.Println(boxTopLeft + strings.Repeat(boxTop, boardWidth) + boxTopRight)

	for i := 0; i < boardHeight; i++ {
		fmt.Print(strings.Repeat(" ", leftPadding))
		fmt.Println(boxLeft + strings.Repeat(" ", boardWidth) + boxRight)
	}

	fmt.Print(strings.Repeat(" ", leftPadding))
	fmt.Println(boxBottomLeft + strings.Repeat(boxBottom, boardWidth) + boxBottomRight)
}

func drawSnake() {
	if tooSmall {
		return
	}

	// move cursor
	fmt.Printf("\x1b[%d;%dH", topPadding+2+snake.pos.Y/2, leftPadding+2+snake.pos.X)
	// set foreground to green
	fmt.Print("\x1b[32m")
	// print snake
	if snake.pos.Y%2 == 0 {
		fmt.Print(blockUp)
	} else {
		fmt.Print(blockDown)
	}
	// reset colors
	fmt.Print("\x1b[0m")
}

func drawFood() {
	if tooSmall {
		return
	}

	// move cursor
	fmt.Printf("\x1b[%d;%dH", topPadding+2+food.pos.Y/2, leftPadding+2+food.pos.X)

	// check if snake head and food occupy the same cell
	if food.pos.X == snake.pos.X && food.pos.Y/2 == snake.pos.Y/2 {
		// set background to green
		fmt.Print("\x1b[42m")
	}

	// set foreground to red
	fmt.Print("\x1b[31m")

	// print apple
	if food.pos.Y%2 == 0 {
		fmt.Print(blockUp)
	} else {
		fmt.Print(blockDown)
	}

	// reset colors
	fmt.Print("\x1b[0m")
}

func drawGameOver() {
	msg := "Game over!"
	// move cursor
	fmt.Printf("\x1b[%d;%dH", rows/2, cols/2-len(msg)/2+1)
	fmt.Print(msg)
}
