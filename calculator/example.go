package main

import "fmt"

type Object struct {
	num1 int
	num2 int
}

// 加法类
type Add struct {
	Object
}

func (p *Add) GetResult(a int, b int) int {
	p.num1 = a
	p.num2 = b
	return p.num1 + p.num2
}

// 减法类
type Sub struct {
	Object
}

func (p *Sub) GetResult(a int, b int) int {
	p.num1 = a
	p.num2 = b
	return p.num1 - p.num2
}

func Main() {
	var add Add
	n := add.GetResult(10, 5)
	fmt.Println(n)

	var sub Sub
	n2 := sub.GetResult(10, 5)
	fmt.Println(n2)
}
