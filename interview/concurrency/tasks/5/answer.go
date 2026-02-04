package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
	Что выведет код? Как исправить?
	Предложить 3 способами
*/

func main() {
	counter := 0
	wg := &sync.WaitGroup{}
	mx := &sync.Mutex{}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			mx.Lock()
			counter++
			mx.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println(counter)
}

func main() {
	counter := atomic.Int32{}
	wg := &sync.WaitGroup{}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			counter.Add(1)
		}()
	}

	wg.Wait()

	fmt.Println(counter.Load())
}

func main() {
	ch := make(chan int)
	counter := 0

	for i := 0; i < 100; i++ {
		go func() {
			ch <- i
		}()
	}

	for range 100 {
		v := <-ch

		counter = v
	}

	fmt.Println(counter)
}
