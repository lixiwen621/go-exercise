package main

import (
	"fmt"
	"io"
	"os"
)

// 文件拷贝，将已有的文件复制一份，同时重新命名
func copyFile() {
	// 1. 打开原有文件(file.txt)
	file, err := os.Open("/Users/lixiwen/GolandProjects/file/test.txt")
	if err != nil {
		fmt.Println(err)
	}
	// 2. 创建一个新的文件
	file2, err := os.Create("/Users/lixiwen/GolandProjects/file/test2.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	defer file2.Close()
	// 3. 将原有文件中的内容读取出来，然后写入到新的文件中
	// 每次读取 1KB 的文件内容
	// 下一次读取的数据会覆盖上一次读取的数据
	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		n2, err := file2.Write(buffer[:n])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(n, n2)
	}
	// 4. 关闭文件
}

func main() {
	copyFile()
}
