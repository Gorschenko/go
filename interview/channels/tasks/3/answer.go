package main

import "fmt"

/*
	Произвести запись из 1 горутины, а прочитать из 2
*/

func main() {
	ch := make(chan int)

	go func() {
		for i := range 10000 {
			ch <- i
		}
	}()

	go func() {
		for v := range ch {
			fmt.Println("v = ", v, "worker = 1")
		}
	}()

	for v := range ch {
		fmt.Println("v = ", v, "worker = 2")
	}

	close(ch)
}
