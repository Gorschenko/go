package main

import "fmt"

/*
	Объяснить построчно, что будет.
	Предложить 2 решения.
*/

func spawnMessages(n int) chan string {
	ch := make(chan string, 1) // создаем канал с буффером

	// пишем в блокирующем цикле, 1 сообщение запишется, так как буфер = 1
	// на 2 сообщение будет блокировка, так как из канала никто не читает:
	// канал еще не вернули
	for i := 0; i < n; i++ {
		ch <- fmt.Sprintf("msg %d", i+1)
	}

	return ch
}

func main() {
	n := 10

	for msg := range spawnMessages(n) { // читаем из канала в цикле
		fmt.Println("received: ", msg) // выводим сообщения
	}
}

func spawnMessages2(n int) chan string {
	ch := make(chan string, 1)

	go func() {
		for i := 0; i < n; i++ {
			ch <- fmt.Sprintf("msg %d", i+1)
		}
		close(ch)
	}()

	return ch
}

func spawnMessages3(n int) chan string {
	ch := make(chan string, n)

	for i := 0; i < n; i++ {
		ch <- fmt.Sprintf("msg %d", i+1)
	}
	close(ch)

	return ch
}
