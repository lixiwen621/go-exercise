package _range

import (
	"fmt"
	"time"
)

// range遍历字符串
func method1() {
	// range遍历字符串，str1[index] 打印的是 Unicode值
	str1 := "abc123"
	for index := range str1 {
		fmt.Println("str1 -- index:%d, value:%d\n", index, str1[index])
	}

	// 跟上面的打印是一样的, value可以直接打印出 index对应的值
	for index, value := range str1 {
		fmt.Println("str1 -- index:%d, value:%d\n", index, value)
	}

	str2 := "测试中文"
	for index := range str2 {
		fmt.Printf("str2 -- index:%d, value:%d\n", index, str2[index])
	}
	fmt.Printf("len(str2) = %d\n", len(str2))

	runesFromStr2 := []rune(str2)
	bytesFromStr2 := []byte(str2)
	fmt.Printf("len(runesFromStr2) = %d\n", len(runesFromStr2))
	fmt.Printf("len(bytesFromStr2) = %d\n", len(bytesFromStr2))
}

// 打印数组和切片
func method2() {
	array := [...]int{1, 2, 3}
	slice := []int{4, 5, 6}

	// 方法1：只拿到数组的下标索引
	for index := range array {
		fmt.Printf("array -- index=%d value=%d \n", index, array[index])
	}
	for index := range slice {
		fmt.Printf("slice -- index=%d value=%d \n", index, slice[index])
	}
	fmt.Println()

	// 方法2：同时拿到数组的下标索引和对应的值
	for index, value := range array {
		fmt.Printf("array -- index=%d index value=%d \n", index, array[index])
		fmt.Printf("array -- index=%d range value=%d \n", index, value)
	}
	for index, value := range slice {
		fmt.Printf("slice -- index=%d index value=%d \n", index, slice[index])
		fmt.Printf("slice -- index=%d range value=%d \n", index, value)
	}
	fmt.Println()
}

// 二维数组使用 range函数
func method3() {
	array := [...][3]int{{1, 2, 3}, {4, 5, 6}}
	slice := [][]int{{1, 2}, {4}}
	// 二维数组遍历为 一维数组
	for index, arrayValue := range array {
		fmt.Printf("array -- index=%d value=%d \n", index, arrayValue)
	}

	for index, sliceValue := range slice {
		fmt.Printf("slice -- index=%d index value=%d \n", index, sliceValue)
	}

	// 把二维数组 遍历为元素value值，也就是把二维数组拆分为value值(双重遍历)
	for row_index, row_value := range array {
		for col_index, col_value := range row_value {
			fmt.Printf("array[%d][%d]=%d ", row_index, col_index, col_value)
		}
		fmt.Println()
	}

	for row_index, row_value := range slice {
		for col_index, col_value := range row_value {
			fmt.Printf("slice[%d][%d]=%d ", row_index, col_index, col_value)
		}
		fmt.Println()
	}

}

// 对通道的迭代
func method4() {
	ch := make(chan int, 10)

	go addData(ch)

	for i := range ch {
		fmt.Println(i)
	}
}

func addData(ch chan int) {
	size := cap(ch)
	for i := 0; i < size; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}
	close(ch)
}

// range 遍历 map
func method5() {
	hash := make(map[string]int)
	hash["abv"] = 1
	hash["b"] = 2
	hash["c"] = 3
	hash["d"] = 4

	for key := range hash {
		fmt.Printf("key=%s, value=%d\n", key, hash[key])
		fmt.Printf("key=%s, value=%d\n", key, hash[key])
	}

	for key, value := range hash {
		fmt.Printf("key=%s, value=%d\n", key, value)
	}
}
