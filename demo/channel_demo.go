package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	ch <- 10

	<-ch
	fmt.Println("ads")
}
