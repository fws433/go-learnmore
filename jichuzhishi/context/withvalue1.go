package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type TraceCode string
var wg4 sync.WaitGroup
func main(){
	//设置一个50ms 的超时
	ctx,cancel:=context.WithTimeout(context.Background(),time.Millisecond*50)
	ctx=context.WithValue(ctx,TraceCode("TRACE_CODE"),"12345")
	wg4.Add(1)
	go worker6(ctx)
	time.Sleep(time.Second*5)
	cancel()
	wg4.Wait()
	fmt.Println("over")
}
func worker6(ctx context.Context) {
	key := TraceCode("TRACE_CODE")
	traceCode, ok := ctx.Value(key).(string) // 在子goroutine中获取trace code
	if !ok {
		fmt.Println("invalid trace code")
	}
LOOP:
	for {
		fmt.Printf("worker, trace code:%s\n", traceCode)
		time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
		select {
		case <-ctx.Done(): // 50毫秒后自动调用
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg4.Done()
}