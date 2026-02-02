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
	totalWorkSeconds := 0
	start := time.Now()

	for range 100 {
		randomWait()
	}

	mainSeconds := time.Since(start)
	fmt.Println("main: ", mainSeconds)
	fmt.Println("total: ", totalWorkSeconds)
}
