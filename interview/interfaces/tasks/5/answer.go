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
	foo(a) // false, так как указан тип
	foo(b) // false, так как указаны данные
	foo(c) // true, так как указан пустой тип
}
