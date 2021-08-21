package main

import "fmt"

// type number struct {
// 	num, num1 int
// }

func main() {

	ch := make(chan int)

	go func() {
		ch <- 75
		ch <- 100
		close(ch)

	}()

	// print pake for
	for v := range ch {
		fmt.Println(v)
	}

	// var num = <-ch
	// var num1 = <-ch

	// fmt.Println(num)
	// fmt.Println(num1)

}
