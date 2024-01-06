package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg = sync.WaitGroup{}
	// ch := make(chan int, 50)
	ch1 := make(chan int, 50)
	ch2 := make(chan int, 50)

	// wg.Add(2)
	// go func(ch <-chan int) {
	// 	for {
	// 		if i, ok := <-ch; ok {
	// 			fmt.Println(i)
	// 		} else {
	// 			break
	// 		}
	// 	}
	// 	wg.Done()
	// }(ch)
	// go func(ch chan<- int) {
	// 	ch <- 42
	// 	ch <- 42
	// 	close(ch)
	// 	wg.Done()
	// }(ch)

	wg.Add(2)
	go func() {
		for {
			select {
			case i := <-ch1:
				fmt.Println("Channel 1: %v", i)
			case i := <-ch2:
				fmt.Println("Channel 2: %v", i)
			default:
				break
			}
		}
		wg.Done()
	}()
	go func() {
		ch1 <- 42
		close(ch1)
		ch2 <- 27
		close(ch2)
		wg.Done()
	}()
	wg.Wait()
}
