package main

import (
	"fmt"
	"strings"
)

var tooSmall bool
var topPadding int
var leftPadding int

func drawBoard() {
	// clear the screen
	fmt.Print("\x1b[H\x1b[2J")

	// account for borders and newline
	tooSmall = rows < boardHeight+3 || cols < boardWidth+2
	if tooSmall {
		fmt.Print("terminal too small")
		return
	}

	topPadding = (rows - boardHeight - 3) / 2
	leftPadding = (cols - boardWidth - 2) / 2

	// draw board
	fmt.Print(strings.Repeat("\n", topPadding))

	fmt.Print(strings.Repeat(" ", leftPadding))
	fmt.Println(boxTopLeft + strings.Repeat(boxTop, boardWidth) + boxTopRight)

	for i := 0; i < boardHeight; i++ {
		fmt.Print(strings.Repeat(" ", leftPadding))
		fmt.Println(boxLeft + strings.Repeat(" ", boardWidth) + boxRight)
	}

	fmt.Print(strings.Repeat(" ", leftPadding))
	fmt.Println(boxBottomLeft + strings.Repeat(boxBottom, boardWidth) + boxBottomRight)

	drawSnake()

	if gameState == stateOver {
		drawGameOver()
	}
}

func drawSnake() {
	if tooSmall {
		return
	}

	// move cursor
	fmt.Printf("\x1b[%d;%dH", topPadding+2+snake.pos.Y/2, leftPadding+2+snake.pos.X)
	// print snake
	if snake.pos.Y%2 == 0 {
		fmt.Print(blockUp)
	} else {
		fmt.Print(blockDown)
	}
}

func drawGameOver() {
	// game over
	if gameState == stateOver {
		msg := "Game over!"
		// move cursor
		fmt.Printf("\x1b[%d;%dH", rows/2, cols/2-len(msg)/2+1)
		fmt.Print(msg)
	}
}
