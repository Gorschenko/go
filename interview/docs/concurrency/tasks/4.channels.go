package main

import (
	"context"
	"fmt"
	"time"
)

/*
	Использовать оператор select и не выйти из следующим образом:
	по чтению из канала, по закрытию канала, по контексту,
	по таймеру, по time.After, по дефолту
*/

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	timer := time.NewTimer(2 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond)
	defer cancel()

	select {
	case v := <-ch1:
		fmt.Println("v = ", v, "from ch1")
	case v := <-ch2:
		fmt.Println("v = ", v, "from ch2")
	case <-ch3:
		fmt.Println("From close ch3")
	case <-timer.C:
		fmt.Println("From custom timer")
	case <-time.After(1 * time.Second):
		fmt.Println("From after timer")
	case <-ctx.Done():
		fmt.Println("From context")
		// default:
		// 	fmt.Println("From default")
	}
}
