package main

import (
	"context"
	"fmt"
	"time"
)

func work(ctx context.Context) {
	select {
	case <-time.After(time.Second):
		fmt.Println("work terminado")
	case <-ctx.Done():
		fmt.Println("work cancelado:", ctx.Err())
	}
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	work(ctx)
}
