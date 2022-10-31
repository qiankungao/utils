package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		go func() {
			var n int
			for {
				dst <- n
				n++
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println("gao:", n)
		if n == 5 {
			cancel()
			break
		}
	}
	time.Sleep(1 * time.Second)
}
