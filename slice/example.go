package slice

import "fmt"

/*
*
声明切片，跟数组有些类似
*/
func method1() {
	// 方式1, 声明并初始化一个空的切片
	var s1 []int = []int{1, 2, 3}
	fmt.Println(s1)

	// 方式2, 类型推导, 并初始化一个空的切片
	var s2 = []int{}
	fmt.Println(s2)

	// 方式3, 与方式2 等价
	s3 := []int{}
	fmt.Println(s3)

	// 方式4, 与方式1、2、3 等价, 可以在大括号中定义切片初始元素
	s4 := []int{1, 2, 3, 4}
	fmt.Println(s4)

	// 方式5，用make() 函数创建切片，创建[]int类型的切片，指定切片初始长度为0
	s5 := make([]int, 0)
	fmt.Println(s5)

	// 方式6，用make() 函数创建切片, 创建[]int类型的切片，初始长度为2，指定容量长度为4
	s6 := make([]int, 2, 4)
	fmt.Println(s6)

	// 方式7，引用一个数组，初始化切片
	a := [5]int{6, 5, 4, 3, 2}
	fmt.Println(a)
	// 从数组下标2开始，直到数组的最后一个元素
	s7 := a[2:]
	fmt.Println(s7)
	// 从数组下标1开始，直到数组下标3的元素，创建一个新的切片
	s8 := a[1:3]
	fmt.Println(s8)
	// 从0到下标2的元素，创建一个新的切片
	s9 := a[:2]
	fmt.Println(s9)
}

// 访问切片
func method2() {
	s1 := []int{5, 4, 3, 2, 1}
	fmt.Println(s1)

	// 下标访问切片
	e1 := s1[0]
	e2 := s1[1]
	e3 := s1[2]
	fmt.Println(e1, e2, e3)

	//向指定位置赋值
	s1[0] = 100
	s1[1] = 10
	s1[2] = 9

	// range 迭代访问切片
	for i, v := range s1 {
		fmt.Println("before modify, s1[%d] = %d", i, v)
	}
}

func method3() {
	var s1 []int = []int{}
	fmt.Println(len(s1))
	fmt.Println(cap(s1))

	s2 := []int{1, 2, 3}
	fmt.Println(len(s2))
	fmt.Println(cap(s2))
}

func method4() {
	s3 := []int{}
	fmt.Println(s3)

	// 使用 append 函数追加元素
	s3 = append(s3, 1)
	s3 = append(s3, 2, 3)
	fmt.Println(s3)

}

// 向指定位置插入元素
func method5() {
	s4 := []int{1, 2, 3, 4}
	// s[:2] 是 s[0] 和 s[1] → [1, 2]
	// s[2:] 是 s[2] 开始往后的部分 → [4, 5]
	s4 = append(s4[:2], append([]int{10, 20}, s4[2:]...)...)
	fmt.Println(s4)
}

// 指定位置移除元素
func method6() {
	s5 := []int{1, 2, 3, 5, 4}
	s5 = append(s5[:3], s5[4:]...)
	fmt.Println(s5)
}

// 复制元素
func method7() {
	src1 := []int{1, 2, 3}
	// 创建一个长度为3, 容量为5的 int数组
	dst1 := make([]int, 3, 5)

	src2 := []int{1, 2, 3, 4, 5}
	dst2 := make([]int, 3, 5)
	// 注意顺序 是dst1 复制 src1 的数据
	copy(dst1, src1)
	copy(dst2, src2)
	fmt.Println(dst1)
	fmt.Println(src1)
	fmt.Println(dst2)
	fmt.Println(src2)
}

// 使用 append 函数未触发 扩容
func method8() {
	s := make([]int, 3, 6)
	fmt.Println("initial, s =", s)
	// initial, s = [0 0 0]
	s[1] = 2
	fmt.Println("after set position 1, s =", s)
	// after set position 1, s = [0 2 0]

	s2 := append(s, 4)
	fmt.Println("after append, s2 length:", len(s2))
	// after append, s2 length: 4
	fmt.Println("after append, s2 capacity:", cap(s2))
	// after append, s2 capacity: 6
	fmt.Println("after append, s =", s)
	// after append, s = [0 2 0]
	fmt.Println("after append, s2 =", s2)
	// after append, s2 = [0 2 0 4]

	s[0] = 1024
	fmt.Println("after set position 0, s =", s)
	// after set position 0, s = [1024 2 0]
	fmt.Println("after set position 0, s2 =", s2)
	// after set position 0, s2 = [1024 2 0 4]

	appendInFunc1(s)
	fmt.Println("after append in funcs, s =", s)
	// after append in funcs, s = [1024 2 512]
	fmt.Println("after append in funcs, s2 =", s2)
	// after append in funcs, s2 = [1024 2 512 1022]
}

func appendInFunc1(param []int) {
	param = append(param, 1022)
	fmt.Println("in funcs, param =", param)
	// in funcs, param = [1024 2 0 1022]
	param[2] = 512
	fmt.Println("set position 2 in funcs, param =", param)
	// set position 2 in funcs, param = [1024 2 512 1022]
}

// 使用 append 函数触发了扩容
func method9() {
	s := make([]int, 2, 2)
	fmt.Println("initial, s =", s)

	s2 := append(s, 4)
	fmt.Println("after append, s length:", len(s))
	// after append, s length: 2
	fmt.Println("after append, s capacity:", cap(s))
	// after append, s capacity: 2
	fmt.Println("after append, s2 length:", len(s2))
	// after append, s2 length: 3
	fmt.Println("after append, s2 capacity:", cap(s2))
	// after append, s2 capacity: 4
	fmt.Println("after append, s =", s)
	// after append, s = [0 0]
	fmt.Println("after append, s2 =", s2)
	// after append, s2 = [0 0 4]

	s[0] = 1024
	fmt.Println("after set position 0, s =", s)
	// after set position 0, s = [1024 0]
	fmt.Println("after set position 0, s2 =", s2)
	// after set position 0, s2 = [0 0 4]

	appendInFunc(s2)
	fmt.Println("after append in funcs, s2 =", s2)
	// after append in funcs, s2 = [0 0 4]
}

func appendInFunc(param []int) {
	param1 := append(param, 511)
	param2 := append(param1, 512)
	fmt.Println("in funcs, param1 =", param1)
	// in funcs, param1 = [0 0 4 511]
	param2[2] = 500
	fmt.Println("set position 2 in funcs, param2 =", param2)
	// set position 2 in funcs, param2 = [0 0 500 511 512]
}

func Method10() {
	s := make([]int, 5, 5)
	Modify(s)
	fmt.Println(s)
}

func Modify(sli []int) {
	for i := 0; i < 5; i++ {
		sli = append(sli, i)
	}
}
