package structure

import "fmt"

type Student struct {
	Person4         // 相当于继承了 Person
	score   float64 // 成绩
}

type Student2 struct {
	*Person4 // 指针类型的匿名字段
	score    float64
}

type Teacher struct {
	Person4
	salary float64 // 工资
}

type Person4 struct {
	id   int
	name string
	age  int
}

func Main() {
	// 全部初始化
	var stu Student = Student{Person4: Person4{101, "张三", 18}, score: 90}
	// 部门初始化
	var stu2 Student = Student{score: 90}
	// 指针类型的初始化 需要用上 &
	var stu3 Student2 = Student2{Person4: &Person4{102, "张三2", 19}, score: 95}

	fmt.Println("张三的成绩: ", stu.score)
	fmt.Println(stu)
	fmt.Println(stu2)
	fmt.Println(stu3)
}
