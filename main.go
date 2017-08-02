package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"unsafe"
)

var rows, cols uint16

func init() {
	rows, cols = getSize()
}

func main() {
	// detect terminal resizes
	go func() {
		fmt.Println(getSize())
		winch := make(chan os.Signal, 1)
		signal.Notify(winch, syscall.SIGWINCH)
		for {
			<-winch
			rows, cols = getSize()
			fmt.Println(rows, cols)
		}
	}()

	select {}
}

func die(err error) {
	if err != nil {
		panic(err)
	}
}

func getSize() (rows uint16, cols uint16) {
	var dimensions [4]uint16
	_, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&dimensions)),
	)
	if errno != 0 {
		panic(errno)
	}
	return dimensions[0], dimensions[1]
}
