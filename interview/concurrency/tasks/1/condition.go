package main

import (
	"fmt"
)

/*
	Как исправить, чтобы все было выведено?
*/

func foo() {
	fmt.Println("Hello from foo")
}

func main() {
	go func() {
		fmt.Println("Hello from goroutine")
	}()

	go foo()

	fmt.Println("Hello from main")
}
