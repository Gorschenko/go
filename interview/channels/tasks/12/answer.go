package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

/*
	Написать predicatableFunc, которая будет работать не более 5 сек
*/

func unpredictableFunc() int {
	n := rand.Intn(10)
	time.Sleep(time.Duration(n) * time.Second)
	return n
}

func predictableFunc(ctx context.Context) (int, error) {
	ch := make(chan int)
	var result int

	var cancel context.CancelFunc
	if _, ok := ctx.Deadline(); !ok {
		ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
	}

	go func() {
		result = unpredictableFunc()
		close(ch)
	}()

	select {
	case <-ch:
		return result, nil
	case <-ctx.Done():
		return 0, errors.New("Timeout")
	}

}

func main() {
	result, _ := predictableFunc(context.Background())
	fmt.Println("result = ", result)
}
