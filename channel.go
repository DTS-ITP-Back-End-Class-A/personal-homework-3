
package main

import (
	"fmt"
)

func main() {
	//go goroutine
	ch := make(chan int)

	go func() {
		ch <- 75
		ch <- 100
		close(ch) 

	}()

	for  v := range ch {
	
		fmt.Println(v)

	}

}
//channel bukan seperti file, kita tidak selalu membutuhkan close. close hanya diperlukan 
//ketika penerima harus memberitau bahwa tidak ada  nilai baru lagi, seperti pada looping
