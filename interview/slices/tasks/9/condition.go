package main

import "fmt"

func getBytes(start, end int) []byte {
	arr := [999999999]byte{} // большой объем данных - около 1 Гб
	slice := arr[start:end]
	return slice // возвращается ссылка
}

//
func main() {
	s := getBytes(10, 20) // работаем с большим слайсом, пока не завершится функция main
	fmt.Println(s)
}
