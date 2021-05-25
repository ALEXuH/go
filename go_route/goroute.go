package main

import (
	"fmt"
	"strings"
	"time"
)

func longWait() {
	fmt.Println("start long wait...")
	time.Sleep(5 * 1e9)
	fmt.Println("end long wait...")
}

func shortWait() {
	fmt.Println("start short wait...")
	time.Sleep(2 * 1e9)
	fmt.Println("end short wait...")
}

func sendData(ch chan string) {
	ch <- "dasd"
	ch <- "dasd1"
	ch <- "dasd2"
	ch <- "dasd3"
	ch <- "dasd4"
	ch <- "dasd5"
}

func getData(ch chan string) {

	for {
		input := <-ch
		fmt.Println(input)
	}
}

// 发送和接受者没有就绪时通道属于阻塞状态
func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

// go协程 并行 主线程结束 协程也结束
// Go 特殊的类型，通道（channel: 数据在通道中进行传递：在任何给定时间，一个数据被设计为只有一个协程可以对其访问，所以不会发生数据竞争。 数据的所有权（可以读写数据的能力）也因此被传递
func main() {
	// fmt.Println("main start")
	// go longWait()
	// go shortWait()
	// time.Sleep(6 * 1e9)
	// fmt.Println("main end")

	// channel
	// ch := make(chan string)
	// go sendData(ch)
	// go getData(ch)
	// time.Sleep(1 * 1e9)

	// 如果容量是0或者未设置，通信仅在收发双方准备好的情况下才可以成功。
	// ch1 := make(chan int)
	// go pump(ch1)
	// fmt.Println(<-ch1)
	//go suck(ch1)
	//time.Sleep(1 * 1e9)

	// 有缓冲通道 消费者 生产者模式

	// 并行for循环
	data := make(map[string]string)
	data["aa"] = "aa"
	data["bb"] = "bb"
	data["cc"] = "cc"
	data["dd"] = "dd"
	for k, v := range data {
		go func(k string, v string) {
			data[k] = strings.Repeat(v, 2)
		}(k, v)
	}
	fmt.Println(data)

	// 信号量模式

}
