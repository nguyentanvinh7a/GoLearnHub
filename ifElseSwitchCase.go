package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"John": 32,
		"Rob":  28,
	}

	if age, isExist := m["Rob"]; isExist {
		fmt.Println(age)
	}

	number := 90
	guess := 70

	if guess < number {
		fmt.Println("Too low")
	} else if guess > number {
		fmt.Println("Too high")
	} else {
		fmt.Println("You got it!")
	}

	number1 := 1
	switch number1 {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two")
		break
		fmt.Println("It's two")
		fallthrough
	default:
		fmt.Println("Not one or two")
	}

	var i interface{} = "Hello"
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
		break
		fmt.Println("This will not print")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("i is another type")
	}
}