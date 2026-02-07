package main

import (
	"fmt"
	"sync"
)

/*
	Записывать данные из 2-х горутин, а читать из другой
*/

func writer() <-chan int {
	ch := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := range 10 {
			ch <- i + 1
		}
	}()

	go func() {
		defer wg.Done()
		for i := range 10 {
			ch <- i + 11
		}
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

func main() {
	ch := writer()

	for v := range ch {
		fmt.Println("v=", v)
	}
}
