package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	j := 0
	for j <= 10 {
		fmt.Println(j)
		j++
	}

	k := 0
	for {
		if k%2 == 0 {
			if k == 4 {
				k++
				continue
			}
			fmt.Println(k)
		}
		k++
		if k == 10 {
			break
		}
	}

	Loop:
		for l := 0; l <= 5; l++ {
			for m := 0; m <= 5; m++ {
				if l > 1 {
					break Loop
				}
				fmt.Println(l * m)
				if m > 3 {
					break
				}
			}
		}

	array := [3]int{1, 2, 3}
	for n := 0; n < len(array); n++ {
		fmt.Println(array[n])
	}
	for index, value := range array {
		fmt.Println(index, value)
	}

	o := map[string]int{
		"Yuh": 19,
		"Tom": 20,
		"Mike": 30,
	}
	for key, value := range o {
		fmt.Println(key, value)
	}

	s := "Hello World"
	for index, value := range s {
		fmt.Println(index, string(value))
	}
}
