package main

import (
	"fmt"
	"sync"
)

/*
	Как исправить? Предложить 2 способа
*/

func main() {
	money := 0
	wg := &sync.WaitGroup{}

	wg.Add(1000)
	for range 1000 {
		go func() {
			defer wg.Done()
			money++
		}()
	}

	wg.Wait()
	fmt.Println(money)
}
