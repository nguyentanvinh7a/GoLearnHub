package main

import "fmt"

func main() {
	// points := []int{20, 30, 40, 50, 60, 70, 80, 90}

	var points [3]string
	points[0] = "Test"
	points[1] = "Test2"
	points[2] = "Test3"
	fmt.Println("%v, %T\n", points[0], points[0])
	fmt.Println("%v, %T\n", points[1], points[1])
	fmt.Println(len(points))
	// fmt.Println(points[2:5])

	a := []int{2, 5, 10, 12, 45, 50}
	b := a[:]
	c := a[3:]
	d := a[:5]
	e := a[3:5]
	e[1] = 100
	fmt.Println("a %v, %T\n", a, a)
	fmt.Println("b %v, %T\n", b, b)
	fmt.Println("c %v, %T\n", c, c)
	fmt.Println("d %v, %T\n", d, d)
	fmt.Println("e %v, %T\n", e, e)

	f := make([]int, 3, 100)
	f = append(f, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	f = append(f, []int{10, 11, 12, 13, 14, 15}...)
	fmt.Println("f %v, %T\n", f, f, len(f), cap(f))

	g := []int{1, 2, 3, 4, 5}
	h := append(g[:2], g[3:]...)
	fmt.Println("h %v, %T\n", h, h)
}