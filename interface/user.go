package main

import "fmt"

func NewUser() {
	// 初始化结构体
	u := User{}
	println(u.Name)
	fmt.Printf("%v \n", u)
	fmt.Printf("%+v \n", u)

	// up 是一个指针
	up := &User{}
	fmt.Printf("up %+v \n", up)
	up2 := new(User)
	fmt.Printf("up2 %+v \n", up2)

	// 下面的 u4和u5的初始化 建议采用 u5的初始化
	u4 := User{Name: "Tom", Age: 18}
	u5 := User{"Tom", 18}

	u4.Name = "Jerry"
	u5.Age = 20

}

type User struct {
	Name string
	Age  int
}

// 这个方法是 值传递
func (u User) changeName(name string) {
	fmt.Printf("changeName中 u的地址 %p \n", &u)
	u.Name = name
}

// 一定要用这个方法接收器才能改变 User对象的值
// 这个方法是引用传递
func (u *User) changeAge(age int) {
	fmt.Printf("changeAge中 u的地址 %p \n", u)
	u.Age = age
}

func ChangeUser() {
	u1 := User{Name: "Tom", Age: 18}
	u1.changeAge(35)

	// 这一步执行的时候，其实相当于复制了一个 u1，改的是复制体
	// 所以 u1 原封不动
	u1.changeName("Jerry")
	fmt.Printf("%+v \n", u1)

	up1 := &User{}
	// 这个也不会生效
	up1.changeName("Jerry")
	up1.changeAge(35)
	fmt.Printf("%+v \n", up1)
}

// Integer 是 int的衍生类型
type Integer int

func UseInt() {
	i1 := 10
	i2 := Integer(i1)
	var i3 Integer = 12
	fmt.Println(i2)
	fmt.Println(i3)
}
