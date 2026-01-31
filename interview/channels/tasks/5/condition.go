package main

import (
	"math/rand"
	"time"
)

/*
	Есть функция, которая выполняется до 100 сек.
	Написать обертку, которая прервет выполнение,
	если она выполняется более 3 сек
*/

func randomTimeWork() {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Second)
}

func predictableTimeWork() {

}

func main() {
	predictableTimeWork()
}
