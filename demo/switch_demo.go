package main

import (
	"fmt"
)

// select  解决channel堵塞问题
func main() {
	ch := make(chan int, 10)
	ch1:= make(chan int, 5)

	for i:=0; i < 10; i++ {
		ch <- i
	}

	for i:=0; i<5; i++ {
		ch1 <- i
	}

	v1 := <- ch
	v2 := <- ch
	fmt.Println(v1, v2)
	//
	//for {
	//	select {
	//	case v3 := <- ch:
	//		fmt.Println("ch channel value", v3)
	//	case v4 := <- ch1:
	//		fmt.Println("ch1 channel value", v4)
	//	default:
	//		fmt.Println(" default value")
	//		return
	//	}
	//}

	//
}
