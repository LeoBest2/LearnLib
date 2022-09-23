package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	echoTime(ctx)

	time.Sleep(time.Second * 200)
}

func echoTime(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("done!")
			fmt.Println("时间到了，可以Ctrl + C 结束了 ^_^ ")
			return
		default:
			time.Sleep(time.Millisecond * 500)
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		}
	}
}
