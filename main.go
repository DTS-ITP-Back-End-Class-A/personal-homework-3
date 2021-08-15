package main

import (
	"fmt"
)

func main() {

	// 1 channel harus 1 tipe data yang sama
	// bisa 2 tipe data tapi pake struct.
	ch := make(chan int)

	go func () {
		ch <- 12
		ch <- 100
		close(ch)
	}()

	for i := 1; i < 3; i++ {
		value := <-ch 
		fmt.Println(value)
	}

}

//mengapa menggunakan goroutine dan channel?
//untuk mengefisiensikan waktu karna concurrency, jadi tidak perlu menunggu(paralel).