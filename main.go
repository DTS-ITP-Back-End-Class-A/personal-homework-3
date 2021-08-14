package main

import (
	"fmt"
	"sync"
)

func worker(wg *sync.WaitGroup, ch chan int, i int) {
	switch i {
	case 1:
		ch <- 75
	case 2:
		ch <- 100
	}
	defer wg.Done()
}

func main() {
	ch := make(chan int, 3)
	wg := &sync.WaitGroup{}

	for i := 1; i < 3; i++ {
		wg.Add(1)
		go worker(wg, ch, i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
