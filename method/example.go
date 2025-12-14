package method

import "fmt"

type Person struct {
	id   int
	name string
	age  int
}

type Student struct {
	Person
	score float64
}

func (p *Person) PrintInfo() {
	fmt.Println("这是父类中的方法")
}

// 这里就是重写了父类的方法
func (s *Student) PrintInfo() {
	fmt.Println("这是子类中的方法")
}

func Main2() {
	stu := Student{Person{101, "张三", 22}, 78}
	stu.PrintInfo()
}
