package main

import "fmt"

// go协程 并行 主线程结束 协程也结束
// Go 特殊的类型，通道（channel: 数据在通道中进行传递：在任何给定时间，一个数据被设计为只有一个协程可以对其访问，所以不会发生数据竞争。 数据的所有权（可以读写数据的能力）也因此被传递
// channel 关闭后再继续读取不会阻塞当读取完时会返回boolean标记，未关闭继续读取会出现阻塞并死锁
// 通过select关键字解决未关闭时阻塞死锁问题
func main() {

	boolChan := make(chan bool)

	intChan := make(chan int, 10)
	go write(intChan)
	go read(intChan, boolChan)
	// 控制主线程阻塞，channel读取完结束程序
	for {
		select {
		case c := <-intChan: // 读取到进入case否则进入下一个case
			fmt.Println("int read 。。。", c)
		case v := <-boolChan:
			fmt.Println("bool read 。。。", v)
			if v {
				return
			}
		default:
			fmt.Println("default..")
		}
	}
}

func read(ch chan int, ch1 chan bool) {
	for {
		v, ok := <-ch
		fmt.Println("value", v, "bool", ok)
		if !ok {
			ch1 <- true
			close(ch1)
			return
		}
	}

}

func write(ch chan int) {
	for i := 0; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}
