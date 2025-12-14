package _interface

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

func Main() {
	var stu Student
	var te Teacher
	var person Person
	var person2 Person
	person = &stu
	person.SayHello()
	person2 = &te
	person2.SayHello()

	whoSayHello(&stu)
	whoSayHello(&te)
}
