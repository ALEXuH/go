package main

import (
	"fmt"
	"io"
	"net"
)

func main()  {
	// 打开服务器监听
	listen,err := net.Listen("tcp", "127.0.0.1:5000")
	if err != nil{
		fmt.Println(" server start fail ", err.Error())
		return
	}
	// 无线循环监听客户端连接　
	fmt.Println("start listing ...")
	for {
		// 阻塞方法
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(" server accept fail", err.Error())
		}
		// 开启协程处理连接
		go func() {
			buf := make([]byte, 20)
			fmt.Println(conn.RemoteAddr().Network(), conn.RemoteAddr().String())
			// 循环读取数据
			for {
				n, err := conn.Read(buf)
				if n != 0 {
					fmt.Println(string(buf))
				}
				if n == 0 && err == io.EOF {
					fmt.Println("data read over")
					// 回复client 关闭连接
					conn.Write([]byte("i am server"))
					return
				}
			}
		}()
	}
}
