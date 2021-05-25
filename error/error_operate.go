package main

import (
	"errors"
	"fmt"
	"strconv"
)

var err = errors.New("error message")

func main() {
	if _, err := errOperate(); err != nil {
		fmt.Println(err)
	}

	// 自定义错误 实现Error接口
	err1 := &SyntaxError{
		msg:    "err1 message",
		offset: 10,
	}

	fmt.Println(err1.Error())
	test()
}

func errOperate() (int, error) {
	return 0, err
}

// 自定义错误 实现Error接口
type SyntaxError struct {
	msg    string
	offset int
}

func (se *SyntaxError) Error() string {
	return se.msg + ":" + strconv.Itoa(se.offset)
}

// 直接使程序崩溃 后续代码不执行
func badCall() {
	panic(" panic error")
}

func test() {
	fmt.Println("start...")
	// recover 恢复程序
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(" panic err ")
		}
	}()
	badCall()
	fmt.Println("end...")
}
