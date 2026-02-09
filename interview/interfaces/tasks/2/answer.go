package main

/*
	Что будет выведено? Объяснить почему.
*/

type ABC interface {
	A()
	B()
	C()
}
type AB interface {
	A()
	B()
}
type BC interface {
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
	var a interface{}
	a = abc{} // v=abc, i=ABC
	//
	ab := a.(AB) // v=abc, i=AB
	ab.A()       // работает
	ab.C()       // ошибка компиляции
	//
	bc := ab.(BC) // v=abc, i=BC
	bc.C()        // работает
	bc.A()        // ошибка компиляции
	//
	abc1 := bc.(ABC) // v=abc, i=ABC
	abc1.A()         // работает
	abc1.B()         // работает
	abc1.C()         // работает
}
