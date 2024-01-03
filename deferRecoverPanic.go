package main

import (
	"fmt"
)

func main() {
	fmt.Println("Line 1")
	defer fmt.Println("Line 2")
	fmt.Println("Line 3")

	a := 1
	defer fmt.Println(a)
	a = 2

	// b := 1
	// c := 0
	// print(b / c)

	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println("Error:", err)
	// 	}
	// }()
	// panic("Panic")

	fmt.Println("Start")
	panicer()
	fmt.Println("End")
}

func panicer() {
	fmt.Println("Creating panic")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
			panic(err)
		}
	}()
	panic("Panic")
}