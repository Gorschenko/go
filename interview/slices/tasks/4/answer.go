package main

func main() {
	a := []int{} // l=0, c=0
	for i := range 3 {
		a = append(a, i+1)
	}
	// i=0, l=1, c=1
	// i=1, l=2, c=2
	// i=2, l=3, c=4, a=[1,2,3],0
	fmt.Println("cap(a) = ", cap(a)) // c=4
	//
	b := append(a, 4) // l=4, c=4, based on a, b=[1,2,3,4]
	c := append(b, 5) // l=5, c=8, new array, c=[1,2,3,4,5],0,0,0
	//
	c[1] = 0 // c=[1,0,3,4,5],0,0,0
	//
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("c =", c)
	//
	d := a[0:4]
	fmt.Prinln("d =", d) // a=[1,2,3,4]
}
