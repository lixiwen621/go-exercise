package main

import "fmt"

type Person interface {
	SayHello()
}

type Student struct {
}

func (stu *Student) SayHello() {
	fmt.Println("student say hello")
}

type Teacher struct {
}

func (te *Teacher) SayHello() {
	fmt.Println("teacher say hello")
}

func whoSayHello(p Person) {
	p.SayHello()
}
