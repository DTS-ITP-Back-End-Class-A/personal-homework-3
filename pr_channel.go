package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 75
		ch <- 100
		close(ch)
	}()

	//bisa pakai for yang ini
	// for i := 1; i <= 2; i++ {
	// 	test := <-ch
	// 	fmt.Println(test)
	// }

	//kalau pakai for range bisa juga
	for v := range ch {
		fmt.Println(v)
		time.Sleep(1 * time.Second)
	}

}
