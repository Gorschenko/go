package main

import "fmt"

/*
	Объяснить построчно, что происходи, а также
	что будет выведено
*/

type c chan c

func main() {
	c := make(c, 1)
	c <- c

	for i := 0; i < 1000; i++ {
		select {
		case <-c:
		case <-c:
			c <- c
		default:
			fmt.Println(i)
			return
		}
	}
}
