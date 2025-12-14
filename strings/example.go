package main

import (
	"fmt"
	"strings"
)

// Contains: 判断一个字符串是否在另一个字符串中
// Join: 字符串连接
// Index: 在一个字符串中查找某个字符串的位置
// Repeat: 某个字符串重复多少次，返回的是重复后的字符串
// Replace: 在 s 字符串中，把 old 字符串替换为 new 字符串，n表示替换的次数，小于0表示全部替换
// Split: 把s字符串按照sep分割，返回slice(切片)

func TestContains() {
	var str string = "hellogo"
	b := strings.Contains(str, "go")
	fmt.Println(b)
}

func TestJoin() {
	s := []string{"abc", "hello", "go"}
	str := strings.Join(s, "|")
	fmt.Println(str)
}

func TestIndex() {
	str := "lixiwen"
	index := strings.Index(str, "xi")
	fmt.Println(index)
}

func TestRepeat() {
	str := "hello"
	result := strings.Repeat(str, 3)
	fmt.Println(result)
}

func TestReplace() {
	str := "lixiwen"
	// 如果n 小于0表示全部替换
	result := strings.Replace(str, "lixiwen", "litongguo", -1)
	fmt.Println(result)
}

func TestSplit() {
	str := "hello_worlds"
	result := strings.Split(str, "_")
	fmt.Println(result)
}

func main() {
	TestSplit()
}
