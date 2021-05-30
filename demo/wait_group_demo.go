package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// 同步并发任务（等待组内协程都结束）
func main() {
	// 初始化计数器 (计数器过多会一直不为0导致死锁)
	wg.Add(3)
	go work()
	go work1()
	go work2()
	// 等待线程都结束（计数器为0）
	wg.Wait()
	fmt.Println("work over")

}

func work() {
	fmt.Println("work")
	defer wg.Done()
}

func work1() {
	fmt.Println("work1")
	defer wg.Done()
}

func work2() {
	fmt.Println("work2")
	defer wg.Done()
}
