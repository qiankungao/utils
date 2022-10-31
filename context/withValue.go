package main

import (
	"context"
	"fmt"
	"time"
)

func main1() {
	ctxVal := make(map[string]string, 0)
	ctxVal["Name"] = "乾坤"
	ctxVal["level"] = "29"
	ctx := context.WithValue(context.Background(), "ctxKey", ctxVal)
	go func(ctx context.Context) {
		data, ok := ctx.Value("ctxKey").(map[string]string)
		if ok {
			fmt.Println("data 钱:", data)
		}
	}(ctx)
	time.Sleep(1 * time.Second)
}
