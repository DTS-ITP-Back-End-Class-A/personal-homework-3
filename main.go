package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 12
		ch <- 100
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}

}
