package main

import "fmt"

func modify(s []int, n int) {
	s = append(s, n) // based on s2, s=[999,2,66],55,0, s1=[999,2,66],55,0, s2=[999,2],66,55,0
	s[0] = 999
}

//
func main() {
	s1 := make([]int, 3, 5) // l=3, c=5, s1=[0,0,0],0,0
	s2 := s1[:2]            // l=2, c=5, based on s1, s2=[0,0],0,0,0
	//
	s1[0] = 1 // s1=[1,0,0],0,0, s2=[1,0],0,0,0
	s2[1] = 2 // s1=[1,2,0],0,0, s2=[1,2],0,0,0
	//
	modify(s1, 55) // s=[999,2,0,55],0, s1=[999,2,0],55,0, s2[]=[999,2],0,55,0
	modify(s2, 66) // s=[999,2,66],55,0, s1=[999,2,66],55,0, s2=[999,2],66,55,0
	//
	fmt.Println("s1 =", s1)
	fmt.Println("s2 = ", s2)
	fmt.Println("s1 cap =", cap(s1))
	fmt.Println("s2 cap =", cap(s2))
	//
	s3 := s2[0:5]
	fmt.Println("s3 =", s3) // s2=[999,2,66,55,0]
}
