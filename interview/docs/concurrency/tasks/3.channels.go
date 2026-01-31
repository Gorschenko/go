package main

import "fmt"

/*
	Писать из одной гоурутины, а читать в нескольких
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
