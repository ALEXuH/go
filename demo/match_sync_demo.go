package main

import (
	"fmt"
	"regexp"
	"sync"
	"time"
)

var (
	mu        sync.Mutex
	valueLock = 10
	value     = 10
)

// regexp
func main() {

	regexpFunction()
	muFunction()

}

//通过使用 sync 包可以解决同一时间只能一个线程访问变量或 map 类型数据的问题。如果这种方式导致程序明显变慢或者引起其他问题，我们要重新思考来通过 goroutines 和 channels 来解决问题，这是在 Go 语言中所提倡用来实现并发的技术
func muFunction() {
	for i := 0; i <= 3000; i++ {
		go updateValueLock(&valueLock)
	}
	for i := 0; i <= 3000; i++ {
		go updateValue(&value)
	}
	time.Sleep(2)
	fmt.Println(valueLock, value)
}

func regexpFunction() {
	value := "dsdsadad dsaadad da123 432"
	pattern := "aad.*"
	b, _ := regexp.MatchString(pattern, value)
	fmt.Println(b)

	re, _ := regexp.Compile(pattern)
	fmt.Println(re.ReplaceAllString(value, "dd"))

}

func updateValueLock(v *int) {
	mu.Lock()
	*v = *v + 1
	mu.Unlock()
}

func updateValue(v *int) {
	*v = *v + 1
}
