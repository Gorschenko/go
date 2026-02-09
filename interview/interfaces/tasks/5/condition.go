package main

import "fmt"

/*
	Что будет выведено?
*/

func foo(a interface{}) {
	fmt.Println(a == nil)
}

func main() {
	var a int
	var b interface{}
	var c interface{}
	b = 1
	foo(a)
	foo(b)
	foo(c)
}
