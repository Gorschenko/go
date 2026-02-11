package main

import "context"

/*
	Реализовать функцию, которая собирает данные из массива каналов
	в один канал. Предусмотреть отмену контекста
*/

func fanIn(ctx context.Context, chans chan []int) chan int {
}
