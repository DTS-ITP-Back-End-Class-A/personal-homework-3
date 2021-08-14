package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 75
		ch <- 100
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}

	tampung1 := <-ch
	tampung2 := <-ch
	fmt.Println(tampung1)
	fmt.Println(tampung2)
}
