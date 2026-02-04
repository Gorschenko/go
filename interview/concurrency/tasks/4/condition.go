package main

import "fmt"

/*
	Что будет выведено? Как исправить?
*/

func main() {
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}
