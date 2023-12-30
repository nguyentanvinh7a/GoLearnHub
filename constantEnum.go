package main

import "fmt"

const (
	_ = 1 << iota * 10
	red 
	yellow
	blue
	black
)

func main() {
	const i int16 = 1
	fmt.Println("%v, %T\n", i, i)

	fmt.Println("%v, %T\n", red, red)
	fmt.Println("%v, %T\n", yellow, yellow)
	fmt.Println("%v, %T\n", blue, blue)
	fmt.Println("%v, %T\n", black, black)
}
// 1. Constant
// 2. Declaration
// 3. Features
// 4. Enum
// 5. iota