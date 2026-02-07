package main

import "fmt"

func main() {
	a := []int{1, 2, 3}              // l= 3 cap =3
	fmt.Println("cap(a) = ", cap(a)) // c=3

	b := append(a, 4) // l=4, cap=6, new [], b=[1,2,3,4],0,0
	c := append(a, 5) // l=4, cap=6, new [], c=[1,2,3,5],0,0
	c[1] = 0          // c=[1,0,3,5],0,0

	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("c =", c)
}
