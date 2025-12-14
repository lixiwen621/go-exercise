package _interface

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

func main() {
	var stu Student2
	var per Personer
	per = &stu
	per.Say()
	per.SayHello()

	var h Humaner
	//h = per
	// 这种是
	per = h.(Personer)
	h.SayHello()
}
