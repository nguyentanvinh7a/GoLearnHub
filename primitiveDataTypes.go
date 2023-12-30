package main

import "fmt"

func main() {
	var a bool
	a = 1 == 2
	fmt.Println("%v, %T\n", a, a)

	var b string = "Hello World"
	var c []byte = []byte(b)
	fmt.Println("%v, %T\n", c, c)

	var d rune = 'á'
	fmt.Println("%v, %T\n", string(d), d)
}
/*
Primitive Data Types
1. Boolean
2. Numerics
	- Integers
		- Signed Integers: int8, int16, int32, int64
		- Unsigned Integers: uint8, uint16, uint32, uint64
	- Floating Point Numbers
		- float32, float64
	- Complex Numbers
		- complex64, complex128
3. Text
	- Strings
		- UTF-8
	- Characters
		- rune
	- Byte
		- uint8
*/