package main

import "fmt"

var i int = 0
var p1 *int
var p2 *string

func test_pointer() {
	x := 10
	// p是指向x的指针
	p := &x
	// p1也是指向 x的指针
	p1 = &x

	fmt.Println(*p) // 解引用，输出 10
	fmt.Println(p)  // 输出 x的地址
	fmt.Println(p1)
}

// 也可以修改 指针 指向 新的值
func method2() {
	x := 10
	p := &x         // p 是指向 x 的指针
	fmt.Println(*p) // 解引用，输出 10

	*p = 20        // 修改 x 的值
	fmt.Println(x) // 输出 20
}

func method3() {
	var p1 *int // 声明一个指针类型变量需要使用星号
	var p2 *string

	i := 1
	s := "Hello"
	// p1 是指向 i 的指针
	p1 = &i // 初始化指针也可以通过另外一个变量
	// p2 指向 s 的地址
	p2 = &s
	// p3 是指向 p2 的指针 (p2 本身是 *string)
	p3 := &p2
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println(*p1)
	fmt.Println(*p2)
	fmt.Println(**p3)
}

func main() {
	method3()
}
