package main

import "fmt"

/*
	Что будет выведено? Как исправить?
*/

func main() {
	ch := make(chan int, 3)

	go func() {
		for range 3 {
			v := <-ch
			fmt.Println("v =", v)
		}
	}()

	ch <- 1
	ch <- 2
	ch <- 3
}
