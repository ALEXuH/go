package main

import (
	"fmt"
	"time"
)

func main() {

	intChan := make(chan int)
	intChan1 := make(chan int)
	go write(intChan)
	go write(intChan1)
	go read(intChan, intChan1)
	time.Sleep(2e9)
}

func read(ch chan int, ch1 chan int) {
	for {
		select {
		case v := <-ch:
			fmt.Println("ch", v)
		case v1 := <-ch1:
			fmt.Println("ch1", v1)
		}
	}

}

func write(ch chan int) {
	for i := 0; i <= 5; i++ {
		ch <- i
	}
}
