package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	// 文件打开
	filename := "D:\\project\\go_learn\\go\\stand\\aa.txt"
	inputFile, err := os.Open(filename)

	if err != nil {
		fmt.Println("read error : %s", err.Error())
		return
	}

	defer inputFile.Close()

	// 行读取
	bufferIo := bufio.NewReader(inputFile)
	for {
		v, err := bufferIo.ReadString('\n')
		fmt.Println("value is:", v)
		if err == io.EOF {
			break
		}
	}

	// 二进制字节缓冲读取
	inputFile1, _ := os.Open(filename)
	defer inputFile1.Close()
	bufferIo = bufio.NewReader(inputFile1)
	buff := make([]byte, 1024)
	for {
		n, err := bufferIo.Read(buff)
		fmt.Println("n....", n, err)
		if n != 0 {
			//fmt.Println("buff value:", buff)
		}
		if n == 0 && err == io.EOF {
			fmt.Println("file read over..")
			break
		}
	}

	// 一次性读取整个文件
	buf, err := ioutil.ReadFile(filename)
	fmt.Println("whole file ", buf)

	// 文件写入 字符串写入 二进制写入 一次性写入
	writer, err := os.OpenFile("bb.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("file write fail. . .")
	}
	defer writer.Close()

	writer.WriteString("aaa")
	bufWrite := bufio.NewWriter(writer)

	bufWrite.WriteString("aaa")
	bufWrite.Write(buf)
	bufWrite.Flush()

	// 拷贝文件
	writer1, _ := os.OpenFile("cc.txt", os.O_CREATE|os.O_WRONLY, 0666)
	src, err := os.Open(filename)
	io.Copy(writer1, src)
	defer writer1.Close()
	defer src.Close()

	// 读取压缩文件
}
