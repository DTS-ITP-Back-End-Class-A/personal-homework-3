package main

import (
	"fmt"
)

func main() {
	// create new channel of type int
	ch := make(chan int)
	// start new anonymous goroutine
	go func() {
		// send value to channel
		ch <- 75
		ch <- 100
		// clear
		close(ch)
	}()
	// read from channe
	for val := range ch {
		fmt.Println(val)
	}

}
