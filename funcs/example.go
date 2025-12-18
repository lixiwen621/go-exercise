package main

import "fmt"

// 声明函数
func custom() {
	fmt.Println("Hello World")
}

type A struct {
	i int
}

// 声明方法
func (a *A) add(v int) int {
	a.i += v
	return a.i
}

func mainExample() {
	// 声明函数变量
	var function1 func(int, int) int

	// 赋值给函数变量
	function1 = func(a, b int) int {
		return a + b
	}

	fmt.Println(function1)
}

// 声明闭包, 示例1
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main2() {
	c1 := counter()
	fmt.Println(c1()) // 1
	fmt.Println(c1()) // 2
	fmt.Println(c1()) // 3

	c2 := counter()
	fmt.Println(c2()) // 1 （重新开始）
}

func DeferExample() {
	defer fmt.Println("第一个defer")
	defer fmt.Println("第二个defer")
	defer fmt.Println("第三个defer")
}

var squart2 func(int) int = func(p int) int {
	p *= p
	return p
}

// 声明闭包示例2
