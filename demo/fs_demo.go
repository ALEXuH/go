package main

import "fmt"

func main() {
	ch := make(chan bool)
	go func() {
		for {
			fmt.Println("d")
		}
	}()
	fmt.Println("ds")
	<-ch
}
