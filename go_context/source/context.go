package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 设置一个过期时间为1秒的context
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	ctxs, cancels := context.WithCancel(context.Background())
	defer cancels()
	fmt.Println(ctx.Deadline())
	fmt.Println(ctxs.Deadline())
	//go handle(ctx, 500*time.Millisecond)
	go handle(ctx, 1500*time.Millisecond)

	select {
	case <-ctx.Done(): // 过期关闭context
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())

	case <-time.After(duration): // select阻塞时间
		fmt.Println("process request with", duration)
	}
}
