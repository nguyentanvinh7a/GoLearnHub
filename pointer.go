package main

import (
	"fmt"
)

func main() {
	// var a int = 12
	// var b *int = &a
	// fmt.Println(a, *b)
	// a = 24
	// fmt.Println(a, *b)
	// *b = 36
	// fmt.Println(a, *b)

	// c := []int{1, 2, 3}
	// d := c
	// fmt.Println(c, d)
	// c[1] = 4
	// fmt.Println(&c[1], &d[1])

	// var a *myStruct
	// fmt.Println(a)
	// a = new(myStruct)
	// fmt.Println(a)
	// a.number = 100
	// fmt.Println(a.number)

	var a = map[string]string{"a": "b", "c": "d"}
	b := a
	fmt.Println(a, b)
	a["a"] = "z"
	fmt.Println(a, b)
}

type myStruct struct {
	number int
}