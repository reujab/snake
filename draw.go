package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func draw() {
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
	drawStats()
}

func drawBox() {
	fmt.Print(strings.Repeat("\n", topPadding))

	fmt.Print(strings.Repeat(" ", leftPadding))
	fmt.Println(boxTopLeft + strings.Repeat(boxTop, boardWidth) + boxTopRight)

	for i := 0; i < boardHeight/2; i++ {
		fmt.Print(strings.Repeat(" ", leftPadding))
		fmt.Println(boxLeft + strings.Repeat(" ", boardWidth) + boxRight)
	}

	fmt.Print(strings.Repeat(" ", leftPadding))
	fmt.Println(boxBottomLeft + strings.Repeat(boxBottom, boardWidth) + boxBottomRight)
}

func drawSnake() {
	// set foreground to white
	fmt.Print("\x1b[37m")

	for i, point := range snake.body {
		// move cursor
		fmt.Printf("\x1b[%d;%dH", topPadding+2+point.Y/2, leftPadding+2+point.X)

		// print point
		var fullBlock bool
		for j := range snake.body {
			if j != i && point.X == snake.body[j].X && snake.body[j].Y/2 == point.Y/2 {
				fullBlock = true
				fmt.Print(block)
				break
			}
		}
		if !fullBlock {
			if point.Y%2 == 0 {
				fmt.Print(blockUp)
			} else {
				fmt.Print(blockDown)
			}
		}
	}

	// reset colors
	fmt.Print("\x1b[0m")
}

func drawFood() {
	// move cursor
	fmt.Printf("\x1b[%d;%dH", topPadding+2+food.pos.Y/2, leftPadding+2+food.pos.X)

	// check if snake and food occupy the same cell
	for _, point := range snake.body {
		if food.pos.X == point.X && food.pos.Y/2 == point.Y/2 {
			// set background to white
			fmt.Print("\x1b[47m")
		}
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

func drawStats() {
	// move cursor
	fmt.Printf("\x1b[%d;%dH", topPadding+3+boardHeight/2, leftPadding+2)

	seconds := time.Since(start).Seconds()
	elapsed, _ := time.ParseDuration(strconv.Itoa(int(seconds)) + "s")
	fmt.Println("Elapsed time:", elapsed)

	fmt.Print(strings.Repeat(" ", leftPadding+1))
	fmt.Print("Length: ", len(snake.body))
}
