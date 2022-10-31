package main

import (
	"context"
	"fmt"
	"time"
)

func main2() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println("gao:", ctx.Err())
	}
}
