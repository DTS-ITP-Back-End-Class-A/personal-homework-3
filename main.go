package main

import "fmt"

func main() {
	channel := make(chan int)

	go func() {
		channel <- 12
		channel <- 100
		close(channel)
	}()

	for i := range channel {
		fmt.Println(i)
	}

	result := <-channel
	fmt.Println(result)
}
