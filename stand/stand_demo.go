package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"runtime"
)

//
var options = &struct {
	config     string
	cpuNumber  int
	autoReload bool
	menProfile string
}{}

var name *string

// 解析命令行参数
func init() {
	name = flag.String("name", "go client", "project name")
	flag.StringVar(&options.config, "config", "aa.ini", "config file name")
	flag.IntVar(&options.cpuNumber, "cpuNumber", 10, "cpuNumber")
	flag.BoolVar(&options.autoReload, "autoReload", false, "autoReload")
	flag.StringVar(&options.menProfile, "menProfile", "menProfile", "menProfile ")
	flag.Parse()
}

func main() {

	// fmt 标准输出，输出错误 返回格式化字符串
	fmt.Println("stand out")
	err := fmt.Errorf("error format %s", "500 error")
	fmt.Printf("%v", err)
	fmt.Println(fmt.Sprintf("hi i am %s", "aa"))
	fmt.Println(runtime.NumCPU())

	// flag 返回都为指针类型 go run stand_demo.go --config config.yml --cpuNumber 50 -name aaaaa --menProfile adsdas --autoReload true
	fmt.Println(name, *name, options, options.config)

	// 键盘读取
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please enter your name")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("wrong reading")
		return
	}

	fmt.Println(fmt.Sprintf("your name is %s", input))

	// 

}
