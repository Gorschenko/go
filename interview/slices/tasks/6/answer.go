package main

import "fmt"

const MAX = 5

//
func main() {
	s := generate() // l=4, c=5, s=[1,2,3,4],0
	mutation(s)     // s=[1,2,3,4],-1
	fmt.Println("s =", s)
	fmt.Println(s[0:MAX]) // s=[1,2,3,4,-1]
}

//
func generate() []int {
	out := make([]int, 0, MAX) // l=0, c=5, out=[],0,0,0,0,0
	for i := 1; i < MAX; i++ {
		out = append(out, i)
	}
	return out
}

//
func mutation(s []int) {
	s = append(s, -1) // l=5, c=5, based on out, s=[1,2,3,4,-1]
}
