package main

import "fmt"

/*
	Что будет выведено? Как исправить?
*/

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			ch <- 1
		}
	}()

	for n := range ch {
		fmt.Println(n)
	}
}
