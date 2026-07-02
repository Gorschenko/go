package main

import (
	"fmt"
	"math/rand"
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
	ch := make(chan int)

	go func() {
		for i := range 3 {
			ch <- i + 1
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
