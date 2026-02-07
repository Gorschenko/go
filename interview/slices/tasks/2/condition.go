package main

import "fmt"

func main() {
	a := []int{}
	a = append(a, []int{1, 2, 3, 4, 5}...)
	fmt.Println("cap(a) = ", cap(a))

	b := append(a, 6)
	c := append(a, 7)
	c[1] = 0

	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("c =", c)
}
