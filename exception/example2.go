package main

import "fmt"

func mayPanic() {
	panic("出现严重错误")
}

func main() {
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println("捕获到 panic:", err)
	//	}
	//}()
	//
	//mayPanic()
	//fmt.Println("程序继续运行")
	TryCatch(
		func() {
			fmt.Println("开始执行")
			panic("出错了！")
		},
		func(r interface{}) {
			fmt.Println("捕获到错误:", r)
		},
	)
	fmt.Println("程序继续运行")
}

// 封装处理成工具类的写法
func TryCatch(try func(), catch func(r interface{})) {
	defer func() {
		if r := recover(); r != nil {
			catch(r)
		}
	}()
	try()
}
