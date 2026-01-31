package main

import (
	"fmt"
	"sync"
)

/*
	Как исправить, чтобы все было выведено?
*/

func foo(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello from foo")
}

func main() {
	wg := &sync.WaitGroup{} // нужно для синхронизации гоурутин

	wg.Add(1) // добавляем гоурутину в группу
	go func() {
		defer wg.Done() // сигнализируем о том, что гоурутина закончилась
		fmt.Println("Hello from goroutine")
	}()

	wg.Add(1)
	go foo(wg)

	wg.Wait() // Ждем выполнение всей группы
	fmt.Println("Hello from main")
}
