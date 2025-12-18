package main

import "fmt"

func main() {
	str := Func0("111")
	fmt.Println(str)
	str1, err1 := Func1(1, 2, 3, "test")
	fmt.Println(str1, err1)
	str2, err2 := Func2(1, 2)
	fmt.Println(str2, err2)
	str3, err3 := Func3(1, 2)
	fmt.Println(str3, err3)
	str4 := Func4()
	fmt.Println(str4)
}
