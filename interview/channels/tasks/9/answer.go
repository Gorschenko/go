package main

import "fmt"

/*
	Что будет выведено? Как исправить?
	Будет блокировка, так как чтобы выйти из range
	нужно закрыть канал
*/

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	for n := range ch {
		fmt.Println(n)
	}
}
