package main

import "fmt"

// 记者: 我叫张三，我的爱好是偷拍，我的年龄是34，我是一个男狗仔。
// 程序员: 我叫孙全，我的年龄是23，我是男生，我的工作年限是3年。
// 1. 定义父类
type Person2 struct {
	name string
	age  int
	sex  string
}

// 2. 给父类添加方法
func (p *Person2) SetValue(userName string, userAge int, userSex string) {
	p.name = userName
	p.age = userAge
	p.sex = userSex
}

// TryRename 值接收者：p 是 Person2 的副本，函数里对字段的修改不会影响外部的嵌入字段。
// 对比 SetValue 的 (*Person2) 指针接收者——那是改「同一块数据」。
func (p Person2) TryRename(newName string) {
	p.name = newName
	p.age = 0
}

// ValueReceiverPractice 演示值接收者 vs 指针接收者对「是否能改原数据」的区别。
func ValueReceiverPractice() {
	var rep Reporter
	rep.SetValue("张三", 20, "男")
	fmt.Println("SetValue 后 name =", rep.name, "age =", rep.age)

	// TryRename 会提升到 Reporter 上调用；接收者是 rep.Person2 的一份拷贝
	rep.TryRename("改名失败示例")
	fmt.Println("TryRename 后 name =", rep.name, "age =", rep.age, "（不变：改的是副本）")
}

// 3. 定义相应的子类
// 记者类
type Reporter struct {
	Person2
	Hobby string // 爱好
}

func (r *Reporter) RepSayHello(hobby string) {
	r.Hobby = hobby
	fmt.Printf("我叫%s，我的爱好是%s，我的年龄是%d，我是一个%s狗仔\n", r.name, r.Hobby, r.age, r.sex)
}

// 程序员类
type Programmer struct {
	Person2
	WorkYear int
}

func (p *Programmer) ProSayHello(workYear int) {
	p.WorkYear = workYear
	fmt.Printf("我叫%s，我的年龄是%d，我是%s，我的工作年限是%d年\n", p.name, p.age, p.sex, p.WorkYear)
}

// 4. 给子类添加相应的方法
// 5. 展示效果
func Main() {
	var rep Reporter
	rep.SetValue("张三", 20, "男")
	rep.RepSayHello("偷拍")

	var pro Programmer
	pro.SetValue("李四", 28, "男")
	pro.ProSayHello(3)

	fmt.Println()
	ValueReceiverPractice()
}
