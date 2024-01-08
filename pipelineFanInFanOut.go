package main

import (
	"fmt"
	"time"
)

func main() {
	// print current time in HH:MM:SS format
	fmt.Println("Current time is: ", time.Now().Format("15:04:05"))
	randomNumbers := []int{}
	for i := 0; i <= 10000000000; i++ {
		randomNumbers = append(randomNumbers, i)
	}
	pipeline := generatePipeline(randomNumbers)

	// // Fan out
	c1 := fanOut(pipeline)
	c2 := fanOut(pipeline)
	c3 := fanOut(pipeline)
	c4 := fanOut(pipeline)
	c5 := fanOut(pipeline)
	c6 := fanOut(pipeline)
	c7 := fanOut(pipeline)
	c8 := fanOut(pipeline)
	c9 := fanOut(pipeline)
	c10 := fanOut(pipeline)
	c11 := fanOut(pipeline)
	c12 := fanOut(pipeline)

	// Fan in
	c := fanIn(c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11, c12)

	// Print the sum of the squares of the random numbers
	sum := 0
	for n := range c {
		sum += n
	}
	// for _, n := range randomNumbers {
	// 	sum += n * n
	// }
	fmt.Println(sum)
	fmt.Println("Current time is: ", time.Now().Format("15:04:05"))
}

func generatePipeline(numbers []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range numbers {
			out <- n
		}
		close(out)
	}()
	return out
}

func fanOut(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func fanIn(inputChannels ...<-chan int) <-chan int {
	in := make(chan int)

	go func() {
		for _, c := range inputChannels {
			for n := range c {
				in <- n
			}
		}
		close(in)
	}()

	return in
}