package main

import "fmt"
import "strconv"

var (
	n int = 10
	m int = 20
	h string = "abc"
)

var Number int = 10

func main() {
	// var i int
	// i = 42
	// var j int = 54
	k := "hello"
	// fmt.Println(i)
	// fmt.Println(j)
	fmt.Println(n)
	fmt.Println(m)
	fmt.Println(h)
	fmt.Printf("%v, %T", k, k)

	helloStringYuh := "hello"
	var numberOfDayPerYear int = 365
	fmt.Println(helloStringYuh)
	fmt.Println(numberOfDayPerYear)

	var i int = 10
	fmt.Printf("%v, %T", i, i)

	fmt.Println()
	var j float32 = float32(i)
	fmt.Printf("%v, %T", j, j)

	fmt.Println()
	var l string = strconv.Itoa(i)
	fmt.Printf("%v, %T", l, l)
} 