package main

import "fmt"

/*
	Что будет выведено и как исправить код?
*/

func test(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int", x)
	case string:
		fmt.Println("string", x)
	case nil:
		fmt.Println("nil", x) // x имеет тип - значит != nil
	case func() int:
		fmt.Println("func", x()) // ошибка компиляции, так тип x = interface{}
		//
		f := x.(func() int)      // позволит вывести
		fmt.Println("func", f()) // ошибка компиляции
	default:
		fmt.Println("unknown type")
	}
}
func main() {
	var x = func() int {
		return 1
	}
	x = nil
	test(x)
}
