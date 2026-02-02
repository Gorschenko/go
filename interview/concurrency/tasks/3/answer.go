package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
	Функция main должна отработать за 5 сек, а randomWait 100 раз.
	2 способа
*/

func randomWait() int {
	workSeconds := rand.Intn(5 + 1)
	time.Sleep(time.Duration(workSeconds) * time.Second)

	return workSeconds
}

func main() {
	wg := &sync.WaitGroup{}
	mx := &sync.Mutex{}
	totalWorkSeconds := 0
	start := time.Now()

	wg.Add(100)
	for range 100 {
		go func() {
			defer wg.Done()
			seconds := randomWait()
			mx.Lock()
			totalWorkSeconds = totalWorkSeconds + seconds
			mx.Unlock()
		}()
	}

	wg.Wait()

	mainSeconds := time.Since(start)
	fmt.Println("main: ", mainSeconds)
	fmt.Println("total: ", totalWorkSeconds)
}

func main() {
	ch := make(chan int)
	totalWorkSeconds := 0
	start := time.Now()

	for range 100 {
		go func() {
			seconds := randomWait()
			ch <- seconds
		}()
	}

	for range 100 {
		totalWorkSeconds += <-ch
	}

	mainSeconds := time.Since(start)
	fmt.Println("main: ", mainSeconds)
	fmt.Println("total: ", totalWorkSeconds)
}
