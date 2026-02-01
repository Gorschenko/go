package main

import "fmt"

/*
	Что будет выведено? Как исправить?
*/

func main() {
	ch := make(chan int)

	go func() {
		for i := range 3 {
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println("v = ", v)
	}
}
