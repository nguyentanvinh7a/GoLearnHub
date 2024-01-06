package main

import (
	"fmt"
)

func main() {
	// a := 10
	// b := 20

	// max := findMax(a, b)
	// fmt.Println(max)

	// sum := findSum(1, 2, 3, 4, 5)
	// fmt.Println(sum)

	// res, error := calDivide(4, 0)
	// fmt.Println(res)
	// fmt.Println(error)

	// func() {
	// 	fmt.Println("Hello")
	// }()

	// for i := 0; i < 5; i++ {
	// 	func(i int){
	// 		fmt.Println(i)
	// 	}(i)
	// }

	g := greeter{
		greeting: "Hello",
		name: "Go",
	}

	g.greet()
	fmt.Println(g.name)

}

func findMax (a, b int) (max int) {
	if a > b {
		max = a
	} else {
		max = b
	}
	return
}

func findSum(a ...int) (sum int) {
	for _, v := range a {
		sum += v
	}
	return
}

func calDivide(a, b int) (int, error) {
	result := 0
	if b == 0 {
		return result, fmt.Errorf("can not divide by zero")
	}
	result = a / b
	return result, nil
}

type greeter struct {
	greeting string
	name string
}

func (g *greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.name = "Vinh"
}