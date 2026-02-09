package main

import "fmt"

/*
	Объяснить, что выведется и почему, предложить исправленный вариант
*/

// Объяснение

type SomeStruct struct{}

func foo() interface{} {
	var result *SomeStruct
	return result
}
func main() {
	res := foo()    // переменная как пустой тип; t = указатель, v = nil
	if res != nil { // true
		fmt.Println("res != nil, rest =", res) // res = nil, так как выводится значение
	}
}

// Исправленный вариант

type SomeInterface interface{}

func foo() SomeInterface {
	var result SomeInterface
	return result
}
func main() {
	res := foo()    // переменная как пустой тип, t = nil, v = nil
	if res != nil { // false
		fmt.Println("res != nil, rest =", res)
	}
}
