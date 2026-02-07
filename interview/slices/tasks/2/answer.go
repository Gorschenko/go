package main

import "fmt"

func main() {
	a := []int{}                           // l=0, c=0
	a = append(a, []int{1, 2, 3, 4, 5}...) // l=5, c=6, a=[1,2,3,4,5],0
	fmt.Println("cap(a) = ", cap(a))       // c=6

	b := append(a, 6) // l=6, c=6, based on a, b=[1,2,3,4,5,6], a=[1,2,3,4,5],6
	c := append(a, 7) // l=6, c=6, based on a, c=[1,2,3,4,5,7], b=[1,2,3,4,5,7], a=[1,2,3,4,5],7
	c[1] = 0          // c=[1,0,3,4,5,7], b=[1,0,3,4,5,7], a=[1,0,3,4,5],7

	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("c =", c)
}
