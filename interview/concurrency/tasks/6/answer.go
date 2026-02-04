package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
	Что будет выведено?
	Как изменится, если изменить GOMAXPROCS

	Будет выведено время > 1 миллисекунд, так как
	планировщик будет вытеснять горутины, а потом ставить их в очередь.
	При добавление кол-ва потоков прирост будет минимальным.
*/

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Millisecond)
}

func main() {
	runtime.GOMAXPROCS(1)
	MAX_TASKS := 10_000
	wg := &sync.WaitGroup{}
	wg.Add(MAX_TASKS)
	start := time.Now()

	for range MAX_TASKS {
		go worker(wg)
	}
	wg.Wait()
	fmt.Println(time.Since(start))
}
