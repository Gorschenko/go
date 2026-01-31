package main

/* Написать 3 функции:
   writer - генерирует числа от 1 до 10, передает их writer
   doubler - умножает числа на 2, передает reader
   reader - читает и выводит на экран
*/

func main() {
	reader(doubler(writer()))
}
