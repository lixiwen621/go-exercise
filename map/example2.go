package main

import (
	"fmt"
)

type User struct {
	Name string
}

// Test 演示 Go 1.22 之前 range 循环的经典陷阱，以及新版本的变化
//
// Go 1.22 之前：for _, u := range 中的 u 是所有迭代共享的同一个变量
// 取 &u 会导致 map 中所有 value 指向同一个地址，最终存的都是最后一个元素
//
// Go 1.22 起：每次迭代都会创建新的循环变量，该陷阱已被修复
func Test() {
	userList := []User{
		{Name: "Tom"},
		{Name: "Jack"},
	}

	// 写法一：&u 直接取循环变量地址
	// Go 1.22 之前：所有 value 指向同一个地址（全变成 Jack）
	// Go 1.22+   ：每次迭代地址不同，结果正确
	m1 := make(map[string]*User)
	for _, u := range userList {
		m1[u.Name] = &u
	}
	fmt.Println("--- 写法一：直接取循环变量地址 ---")
	for name, user := range m1 {
		fmt.Printf("key=%s, user=%+v, addr=%p\n", name, *user, user)
	}

	fmt.Println()

	// 写法二：用索引取地址（始终正确，兼容所有 Go 版本）
	m2 := make(map[string]*User)
	for i := range userList {
		m2[userList[i].Name] = &userList[i]
	}
	fmt.Println("--- 写法二：索引取地址（推荐） ---")
	for name, user := range m2 {
		fmt.Printf("key=%s, user=%+v, addr=%p\n", name, *user, user)
	}

	fmt.Println()

	// 写法三：值拷贝后再取地址（Go 1.22 之前需要这样做）
	m3 := make(map[string]*User)
	for _, u := range userList {
		u := u // Go 1.22+ 中这行已经不再需要，但保留无害
		m3[u.Name] = &u
	}
	fmt.Println("--- 写法三：循环体内重新声明 ---")
	for name, user := range m3 {
		fmt.Printf("key=%s, user=%+v, addr=%p\n", name, *user, user)
	}
}
