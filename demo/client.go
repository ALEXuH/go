package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:5000")
	if err != nil {
		fmt.Println("connect refuse connect。。")
		return
	}
	conn.Write([]byte("say hello 11"))
	//conn.Write([]byte("say hello 22"))
	fmt.Println(conn.LocalAddr())
	conn.Close()
}

type FakerStruct interface{}

type Values interface{}

type Column struct {
	ColumnType string
	values     []Values
	fakerRange []interface{}
}
