package main

import (
	"fmt"
)

func main() {
	var ch = make(chan int)

	go func() {
		ch <- 12
		ch <- 100
		close(ch)
	}()

	for i := range ch {
		fmt.Println(i)
	}

	hasil := <-ch
	fmt.Println(hasil)
}
