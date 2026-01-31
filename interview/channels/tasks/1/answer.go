package main

import (
	"fmt"
	"sync"
	"time"
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
		for i := range 5 {
			ch <- i + 1
		}
	}()

	go func() {
		defer wg.Done()
		for i := range 5 {
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

	for {
		val, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("v=", val)
	}
	time.Sleep(1 * time.Second)
}
