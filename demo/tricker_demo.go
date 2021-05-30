package main

import (
	"fmt"
	"time"
)

func main() {
	// 响应一次
	time3 := time.NewTimer(time.Second * 2)
	<-time3.C
	// 定时器 每隔固定时间向channel发送数据
	time2 := time.NewTicker(time.Second * 2)
	time1 := time.NewTicker(time.Second * 1)
	ch := make(chan bool, 1)
	count := 0
	// 无阻塞循环取 循环5次给channel发送标识信号结束程序
	for {
		select {
		case a := <-time1.C:
			fmt.Println("time1 tricker %v", a)
		case b := <-time2.C:
			fmt.Println("time2 tricker %v", b)
			fmt.Println(count)
			//time2.Stop()
			count++
			if count == 5 {
				ch <- true
			}
		case <-ch:
			return
		}
	}
}
