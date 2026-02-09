package main

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
		fmt.Println("nil", x)
	case func() int:
		fmt.Prinln("func", x())
	default:
		fmt.Prinln("unknown type")
	}
}
func main() {
	var x = func() int {
		return 1
	}
	x = nil
	test(x)
}
