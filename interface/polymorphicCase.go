package _interface

import "fmt"

type Stronger interface {
	Read()
	Writer()
}

// 移动硬盘
type MDisk struct {
}

func (m *MDisk) Read() {
	fmt.Println("移动硬盘读取数据")
}

func (m *MDisk) Writer() {
	fmt.Println("移动硬盘写入数据")
}

// U盘
type UDisk struct {
}

func (u *UDisk) Read() {
	fmt.Println("U盘读取数据")
}

func (u *UDisk) Writer() {
	fmt.Println("U盘写入数据")
}

// 定义一个函数
func Computer(c Stronger) {
	c.Read()
	c.Writer()
}

func Main2() {
	var uDisk UDisk
	var mDisk MDisk
	Computer(&uDisk)
	Computer(&mDisk)
}
