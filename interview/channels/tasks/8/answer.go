package main

import (
	"fmt"
	"time"
)

/*
	Что будет выведено в двух случаях и почему?
	1 - отработает моментально,
	так как двум переменным присваивается толкьо канал

	2 - отработает за 2 сек., так ка происходит чтение из канала,
	которое разблокируется только после закрытия канала или прочтения
*/

func worker() <-chan int {
	ch := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		close(ch)
	}()

	return ch
}

func main() {
	start := time.Now()
	_, _ = worker(), worker()

	fmt.Println(time.Since(start))
}

func main() {
	start := time.Now()
	_, _ = <-worker(), <-worker() // последовательное выполнение

	fmt.Println(time.Since(start))
}
