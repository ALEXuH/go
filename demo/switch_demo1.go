package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	Total = 100
	L     = sync.Mutex{}
	sec   = 1
)

func main() {

	//go addTotal()
	//go addTotal()
	//go addTotal()
	//fmt.Println(Total)
	//
	//for {
	//	select {
	//	case <- time.After(time.Second * 2): // 等待2秒结束程序  定时器
	//		sec += 1
	//		fmt.Println(Total, Total / sec)  // 统计速度
	//		return
	//	}
	//}

	tricker := time.NewTicker(time.Second * 2)
	go func() {
		for t := range tricker.C {
			fmt.Println(t)
		}
	}()
	time.Sleep(10)

}

func addTotal() {
	L.Lock()
	for i := 0; i < 10; i++ {
		Total = Total + 1
	}
	L.Unlock()
}
