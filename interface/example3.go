package main

import "fmt"

type Humaner interface {
	SayHello()
}

type Personer interface {
	Humaner
	Say()
}

type Student2 struct {
}

func (stu *Student2) SayHello() {
	fmt.Println("大家好")
}

func (stu *Student2) Say() {
	fmt.Println("你好")
}
