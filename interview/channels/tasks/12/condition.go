package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

/*
	Написать predicatableFunc, которая будет работать не более 5 сек
*/

func unpredictableFunc() int {
	n := rand.Intn(40)
	time.Sleep(time.Duration(n) * time.Second)
	return n
}

func predictableFunc() (int, error) {

}

func main() {
	result, _ := predictableFunc(context.Background())
	fmt.Println("result = ", result)
}
