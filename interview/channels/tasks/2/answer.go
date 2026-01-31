package main

import (
	"fmt"
	"time"
)

/* Написать 3 функции:
   writer - генерирует числа от 1 до 10, передает их writer
   doubler - умножает числа на 2, передает reader
   reader - читает и выводит на экран
*/

func writer() <-chan int {
	ch := make(chan int)

	go func() {
		for i := range 10 {
			ch <- i + 1
		}

		close(ch)
	}()

	return ch
}

func doubler(inputCh <-chan int) <-chan int {
	ch := make(chan int)

	go func() {
		for v := range inputCh {
			time.Sleep(500 * time.Millisecond)
			ch <- v * 2
		}

		close(ch)
	}()

	return ch
}

func reader(inputChan <-chan int) {
	go func() {
		for v := range inputChan {
			fmt.Println(v)
		}
	}()
}

func main() {
	reader(doubler(writer()))
}
