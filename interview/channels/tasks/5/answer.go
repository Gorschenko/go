package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	Есть функция, которая выполняется до 100 сек.
	Написать обертку, которая прервет выполнение,
	если она выполняется более 3 сек
*/

func randomTimeWork() {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Second)
}

func predictableTimeWork() {
	ch := make(chan int)

	go func() {
		randomTimeWork()
		close(ch)
	}()

	select {
	case <-ch:
		fmt.Println("From func")
	case <-time.After(3 * time.Second):
		fmt.Println("Stop")
	}
}

func main() {
	predictableTimeWork()
}
