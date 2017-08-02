package main

import (
	"fmt"
	"strings"
)

func draw() {
	// clear the screen
	fmt.Print("\x1b[H\x1b[2J")

	// account for borders and newline
	if rows < boardHeight+3 || cols < boardWidth+2 {
		fmt.Println("terminal too small")
		return
	}

	topPadding := (rows - boardHeight - 3) / 2
	leftPadding := (cols - boardWidth - 2) / 2

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

	// draw snake
	// move cursor
	fmt.Printf("\x1b[%d;%dH", topPadding+2+snake.pos.Y/2, leftPadding+2+snake.pos.X)
	// print snake
	if snake.pos.Y%2 == 0 {
		fmt.Print(blockUp)
	} else {
		fmt.Print(blockDown)
	}

	if gameState == stateOver {
		msg := "Game over!"
		// move cursor
		fmt.Printf("\x1b[%d;%dH", rows/2, cols/2-len(msg)/2)
		fmt.Print(msg)
	}
}
