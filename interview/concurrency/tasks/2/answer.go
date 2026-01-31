package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
	Как исправить? Предложить 2 способа
*/

func main() {
	// Гонка. Исправленный вариант. Атомики
	var money atomic.Int32
	wg := &sync.WaitGroup{}

	wg.Add(1000)
	for range 1000 {
		go func() {
			defer wg.Done()
			money.Add(1)
		}()
	}

	wg.Wait()
	fmt.Println(money.Load())
}

func main() {
	// Гонка. Исправленный вариант. Мютекс
	mx := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	wg.Add(1000)
	for range 1000 {
		go func() {
			defer wg.Done()
			mx.Lock()
			money++
			mx.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(money)
}
