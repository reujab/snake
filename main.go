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
	defer func() {
		print("\x1b[?1049l") // tput rmcup
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
