package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	// var wg = sync.WaitGroup{}
	// wg.Add(2)
	// go func() {
	// 	count("sheep")
	// 	wg.Done()
	// }()

	// go func() {
	// 	count("fish")
	// 	wg.Done()
	// }()
	// wg.Wait()
	// fmt.Println("done")

	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func count(name string) {
	for i := 0; i < 5; i++ {
		fmt.Println(name, ":", i)
		time.Sleep(time.Second)
	}
}

func sayHello() {
	fmt.Println("hello", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
