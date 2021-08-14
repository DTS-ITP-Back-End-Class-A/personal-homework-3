package main

import "fmt"

func main() {

	ch := make(chan string, 2)
	go func() {
		ch <- "Dua Belas"
		ch <- "Seratus"
		close(ch)
	}()

	for items := range ch {
		fmt.Println(items)
	}
}
