package main

import (
	"fmt"
	"sync"
)

/*
	Произвести запись из 1 горутины, а прочитать из 2
*/

func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println("value = ", v)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println("value = ", v)
		}
	}()

	for i := range 10 {
		ch <- i + 1
	}
	close(ch)

	wg.Wait()
}
