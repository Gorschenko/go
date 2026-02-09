package main

import "fmt"

/*
	Объяснить, что выведется и почему, предложить исправленные вариант
*/

type SomeStruct struct{}

func foo() interface{} {
	var result *SomeStruct
	return result
}
func main() {
	res := foo()
	if res != nil {
		fmt.Println("res != nil, rest =", res)
	}
}
