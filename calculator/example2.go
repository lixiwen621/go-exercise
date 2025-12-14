package main

import (
	"fmt"
)

// 1. 第一个类
type Object2 struct {
	numA int
	numB int
}

type Resulter interface {
	GetResult() int
	SetData(data ...interface{}) bool //对设置的参数进行校验
}

// 2. 加法类
type Add2 struct {
	Object2
}

func (add2 *Add2) GetResult() int {
	return add2.numA + add2.numB
}

func (add2 *Add2) SetData(data ...interface{}) bool {
	var b bool = true
	if len(data) > 2 {
		fmt.Println("参数个数错误!!!")
		b = false
	}

	value, ok := data[0].(int)
	if !ok {
		fmt.Println("第一个数据类型错误")
		b = false
	}

	value1, ok1 := data[1].(int)
	if !ok1 {
		fmt.Println("第二个数据类型错误")
		b = false
	}
	add2.numA = value
	add2.numB = value1
	return b
}

// 3. 减法类
type Sub2 struct {
	Object2
}

func (sub2 *Sub2) GetResult() int {
	return sub2.numA - sub2.numB
}

func (sub2 *Sub2) SetData(data ...interface{}) bool {
	var b bool = true
	if len(data) > 2 {
		fmt.Println("参数个数错误!!!")
		b = false
	}

	value, ok := data[0].(int)
	if !ok {
		fmt.Println("第一个数据类型错误")
		b = false
	}

	value1, ok1 := data[1].(int)
	if !ok1 {
		fmt.Println("第二个数据类型错误")
		b = false
	}
	sub2.numA = value
	sub2.numB = value1
	return b
}

// 4. 乘法类
type Mul2 struct {
	Object2
}

func (mul2 *Mul2) GetResult() int {
	return mul2.numA * mul2.numB
}

func (mul2 *Mul2) SetData(data ...interface{}) bool {
	var b bool = true
	if len(data) > 2 {
		fmt.Println("参数个数错误!!!")
		b = false
	}

	value, ok := data[0].(int)
	if !ok {
		fmt.Println("第一个数据类型错误")
		b = false
	}

	value1, ok1 := data[1].(int)
	if !ok1 {
		fmt.Println("第二个数据类型错误")
		b = false
	}
	mul2.numA = value
	mul2.numB = value1
	return b
}

// 5. 除法类
type Div2 struct {
	Object2
}

func (div2 *Div2) GetResult() int {
	return div2.numA / div2.numB
}

func (div2 *Div2) SetData(data ...interface{}) bool {
	var b bool = true
	if len(data) > 2 {
		fmt.Println("参数个数错误!!!")
		b = false
	}

	value, ok := data[0].(int)
	if !ok {
		fmt.Println("第一个数据类型错误")
		b = false
	}

	value1, ok1 := data[1].(int)
	if !ok1 {
		fmt.Println("第二个数据类型错误")
		b = false
	}
	if value1 == 0 {
		fmt.Println("除数不能为0")
		b = false
	}
	div2.numA = value
	div2.numB = value1
	return b
}

// 对象问题
// 1: 定义一个新的类
type OperatorFactory struct {
}

// 2: 创建一个方法，在该方法中完成对象的创建
func (o *OperatorFactory) CreateOperator(op string) Resulter {
	switch op {
	case "+":
		//add := Add2{Object2{numA, numB}}
		add := new(Add2)
		return add
	case "-":
		sub := new(Sub2)
		return sub
	case "/":
		div := new(Div2)
		return div
	case "*":
		mul := new(Mul2)
		return mul
	default:
		return nil
	}
}

func OperatorWho(r Resulter) int {
	return r.GetResult()
}

func main() {
	var operator OperatorFactory
	obj := operator.CreateOperator("/")
	b := obj.SetData(10, 0)
	if b {
		num := OperatorWho(obj)
		fmt.Println(num)
	}
}
