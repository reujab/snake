package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// switch to alternate terminal screen
	print("\x1b[?1049h") // tput smcup
	// hide the cursor
	print("\x1b[?25l") // tput civis
	defer func() {
		print("\x1b[?1049l") // tput rmcup
		print("\x1b[?25h")   // tput cvvis
	}()

	go func() {
		for {
			tick()
			draw()
			time.Sleep(time.Millisecond * 250)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interrupt
}

func die(err error) {
	if err != nil {
		panic(err)
	}
}
