package main

import (
	"fmt"
	"io"
	"os"
)

func createFile() {
	// 创建文件，需要指定文件的存储路径以及文件名称，注意 "/"
	file, err := os.Create("/Users/lixiwen/GolandProjects/file/test.txt")
	// 判断是否出现异常
	if err != nil {
		fmt.Println(err)
		//file.Close()
	}
	defer file.Close()
	// 对创建的文件进行相关的操作
	// 写入数据
	n, err := file.WriteString("Hello World")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
	// 关闭
	//file.Close()
}

func writeFile() {
	file, err := os.Create("/Users/lixiwen/GolandProjects/file/test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	str := "Hello World"
	bytes := []byte(str)
	n, err := file.Write(bytes)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}

func writeAt() {
	file, err := os.Create("/Users/lixiwen/GolandProjects/file/test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.WriteString("Hello World")
	str := "lixiwen"
	num, _ := file.Seek(0, io.SeekEnd) // 通过该方法可以将光标定位到文件中原有内容的后面, 返回文件中原有数据的长度
	fmt.Println(num)
	bytes := []byte(str)
	// WriteAt() 方法可以决定写入的位置
	n, err := file.WriteAt(bytes, num)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}

func openFile() {
	// O_RDWR -- 该模式是在文本开头中写入数据
	// O_APPEND -- 是追加模式
	file, err := os.OpenFile("/Users/lixiwen/GolandProjects/file/test.txt", os.O_RDWR, 6)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	n, err := file.WriteString("aaaa")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}

func readFile() {
	// 1. 打开要读取的文件
	file, err := os.Open("/Users/lixiwen/GolandProjects/file/test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	// 2. 进行文件内容读取
	// 定义一个字符类型切片，存储从文件读取的数据
	buffer := make([]byte, 1024*2) // 大小为 2KB
	n, err := file.Read(buffer)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(buffer)
	fmt.Println(string(buffer))
	fmt.Println("读取数据的长度", n)
	// 3. 关闭文件
}

// 循环读取文件
func rangeReadFile() {
	file, err := os.Open("/Users/lixiwen/GolandProjects/file/test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	buffer := make([]byte, 10)
	for {
		n, err := file.Read(buffer)
		// 表示达到文件末尾了
		if err == io.EOF {
			break
		}
		fmt.Println(string(buffer[:n]))
	}
}

func Main() {
	createFile()
}
