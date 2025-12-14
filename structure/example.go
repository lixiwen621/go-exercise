package structure

import (
	"fmt"
	"sync"
)

// 定义结构体
type Person struct {
	Name         string            // 字符串类型: 姓名
	Age          int               // int 类型: 年龄
	Class, Grade int               // int 类型, 多个字段定义int类型
	Call         func() byte       // 函数类型: 返回一个byte函数
	Map          map[string]string // 映射函数: string 到 string 的哈希表
	Ch           chan string       // 通道类型: 传输 string 的 channel
	Arr          [32]uint8         // 数组类型: 固定长度为 32的 uint8 数组
	Slice        []interface{}     // 切片类型: 存放任意类型(空接口)的动态数组
	Ptr          *int              // 指针类型: 指向 int的指针
	once         sync.Once         // 并发原语: 用于只执行一次的操作 (常用于懒加载)

}

type Persion3 struct {
	Name string
	Age  int
}

// 定义匿名字段
// 结构体中的字段不是一定要有字段名，也可以仅定义类型
type Custom struct {
	int
	string
	Other string
}

// 定义匿名结构体
// 打印出来的值 就是name=Tom, age=20
var p = struct {
	Name string
	Age  int
}{
	Name: "Tom",
	Age:  20,
}

// 下面这种定义匿名结构体, 只能在函数中定义
func main() {
	p2 := struct {
		Name string
		Age  int
	}{
		Name: "Tom",
		Age:  20,
	}
	fmt.Println(p2)
	// 使用字面量方式初始化
	p3 := Persion3{"Alice", 25}

	// 使用字段名初始化（推荐）
	p4 := Persion3{
		Name: "John",
		Age:  30,
	}
	fmt.Println(p3, p4)
}

// 嵌套结构体的定义
type Address struct {
	City    string
	ZipCode string
}

type Person2 struct {
	Name    string
	Age     int
	Address Address // 嵌套结构体作为字段
}

// 定义结构体的方法
// 方法: 值接受者(不会修改结构体本身,就是不能修改变量本身的值)
// 请注意 Persion3 是函数的接收者类型，说明这个方法属于Persion3的类型
func (p Persion3) SayHello() {
	fmt.Println("Hello, my name is", p.Name)
}

// 方法: 指针接受者(可以修改结构体, 可以修改变量本身的值)
func (p *Persion3) Birthday() {
	p.Age++
}
