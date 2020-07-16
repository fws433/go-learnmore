package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg5 sync.WaitGroup
func main(){
	ctx,cancel:=context.WithTimeout(context.Background(),time.Microsecond*50)
	wg5.Add(1)
	go worker5(ctx)
	time.Sleep(time.Second*5)
	cancel()
	wg5.Wait()
	fmt.Println("over")
}

func worker5(ctx context.Context) {
	LOOP:
		for{
			fmt.Println("db connecting...")
			time.Sleep(time.Millisecond*10)//假设正常连接数据库耗时10ms
			select {
			case<-ctx.Done():// 50ms之后自动调用
				break LOOP
			default:

			}
		}
		fmt.Println("worker5 done")
	wg5.Done()
}
