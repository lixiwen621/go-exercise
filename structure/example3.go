package structure

import "fmt"

// 面向对象的练习
// 定义一个学生类，有六个属性，分别为姓名、性别、年龄、语文、数学、英语成绩
// 第一个方法: 打招呼的方法: 介绍自己叫XX，今年几岁了。是男同学还是女同学。
// 第二个方法: 计算总分与平均分的方法

// 1. 定义结构体
type StudentInfo struct {
	name    string
	sex     string
	age     int
	chinese float64
	math    float64
	english float64
}

// 2. 添加方法
func (s *StudentInfo) SayHello(username string, userAge int, userSex string) {
	// 初始化2.1
	s.name = username
	s.age = userAge
	s.sex = userSex
	// 2.2 初始化后进行值判断
	if s.sex != "男" && s.sex != "女" {
		s.sex = "男"
	}
	if s.age < 1 || s.age > 100 {
		s.age = 18
	}

	fmt.Printf("我叫%s, 年龄是%d, 性别是%s\n", s.name, s.age, s.sex)
}

func (s *StudentInfo) getScore(chinese float64, math float64, english float64) {
	// 1:初始化
	s.chinese = chinese
	s.math = math
	s.english = english
	// 2:进行计算
	sum := s.chinese + s.math + s.english
	// 3:打印输出结果
	fmt.Print("我叫%s, 总分%f, 平均分%f", s.name, sum, sum/3)

}

func Example3Main() {
	var stu StudentInfo
	stu.SayHello("张三", 18, "男")
	stu.getScore(90, 88, 92)
}
