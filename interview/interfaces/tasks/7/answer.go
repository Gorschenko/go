package main

import (
	"errors"
	"fmt"
)

/*
	Объяснить, что выведется и почему, предложить исправленные вариант
*/

// Объяснение

type SomeError struct{}

func (s SomeError) Erorr() string {
	return "some error"
}

func foo() error {
	var result *SomeError
	return result
}

func main() {
	result := foo()    // t=указатель, v=nil, так как error - это интерфейс
	if result != nil { // true
		fmt.Println("Error", result) // nil
		return
	}
}

// Исправленный вариант

type SomeError struct{}

func (s SomeError) Erorr() string {
	return "some error"
}

func foo() error {
	var result error
	return result
}

func main() {
	result := foo()    // t=error, v=Error<nil>
	if result != nil { // true
		fmt.Println("Error", result)
		return
	}
}

// Исправленный вариант

var SomeError = errors.New("some error")

func foo() error {
	var result error
	result = SomeError
	return result
}

//
func main() {
	result := foo()    // t=error, v=SomeError
	if result != nil { // true
		fmt.Println("Error", result)
		return
	}
}
