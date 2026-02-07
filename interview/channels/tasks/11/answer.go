package main

import "fmt"

/*
	Что будет выведено?
*/

func main() {
	ch := make(chan int, 1)

	for i := 0; i < 5; i++ {
		select {
		case val := <-ch:
			fmt.Println(val)
		case ch <- i:
		}
	}
}

// i = 0; запись 0
// i = 1; чтение 0
// i = 2; запись 2
// i = 3; чтение 2
// i = 4; запись 4
