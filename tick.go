package main

func tick() {
	// update snake position
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
}
