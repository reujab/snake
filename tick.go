package main

func tick() {
	if gameState != stateRunning {
		return
	}

	// update snake position
	lastPos := snake.pos
	switch snake.dir {
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
	}
}
