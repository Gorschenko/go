package main

import "fmt"

func main() {
	a := []int{}
	a = append(a, []int{1, 2, 3, 4, 5}...)
	fmt.Println("cap(a) = ", cap(a))

	b := append(a, 6)
	c := append(b, 7)
	c[1] = 0
	d := c[0:12]

	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("c =", c)
	fmt.Println("d =", d)
}
