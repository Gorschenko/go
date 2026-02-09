package main

import "fmt"

/*
	Что будет выведено? Объяснить почему.
*/

type ABC interface {
	A()
	B()
	C()
}
type abc struct{}

func (a abc) A() {}
func (a abc) B() {}
func (a abc) C() {}

type ab struct{}

func (a ab) A() {}
func (a ab) B() {}

func main() {
	var a = abc{}
	a1 := a.(ABC) // так делать нельзя
	fmt.Println(a1)

	var a interface{}
	a = abc{}
	a1 := a.(ABC) // сработает
	fmt.Println(a1)

	var a interface{}
	a = ab{}
	a1 := a.(ABC) // ошибка runtime из-за рефлексии
	fmt.Println(a1)

	var a interface{}
	a = ab{}
	a1 := a.(ABC) // будет работать
	if !ok {
		fmt.Println("not ok")
	}
	fmt.Println(a1)
}
