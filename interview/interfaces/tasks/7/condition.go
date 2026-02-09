package main

import "fmt"

/*
	Объяснить, что выведется и почему, предложить исправленные вариант
*/

type SomeError struct{}

func (s SomeError) Erorr() string {
	return "some error"
}

func foo() error {
	var result *SomeError
	return result
}

func main() {
	result := foo()
	if result != nil {
		fmt.Println("Error", result)
		return
	}
}
