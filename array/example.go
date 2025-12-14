package array

import "fmt"

/*
*
声明数组
*/
func example() {
	// 仅声明
	var a [5]int
	fmt.Println(a)

	// 声明一个 key为int类型, value为string类型的map
	var marr map[int]string
	fmt.Println(marr)
	// map 的 put操作
	marr[1] = "test"

	// 声明以及初始化
	var b [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	//类似 推导的声明方式
	var c = [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(c)

	d := [3]int{1, 2, 3}
	fmt.Println(d)

	// 使用 ... 代替数组长度, ... 代表可变数组长度
	autoLen := [...]string{"auto1", "auto2", "auto3"}
	fmt.Println(autoLen)

	// 声明时初始化指定下标的元素值
	// 例如 下面的 声明下标 0,1,2
	positionInit := [5]string{0: "position1", 1: "position2", 2: "position3"}
	fmt.Println(positionInit)

	// 初始化时, 元素个数不能超过数组声明长度
	//overLen := [2]int{1,2,3}

}

/*
*
读取数组数据
*/
func method2() {
	a := [...]int{1, 2, 3, 4, 5}

	// 方式1：使用下标读取数组
	element := a[2]
	fmt.Println(element)

	// 方式2: 使用 range 遍历
	for i, v := range a {
		fmt.Println(i, v)
	}

	for i := range a {
		fmt.Println(i)
	}
}

func method3() {
	// 二维数组
	a := [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(a)

	// 三维数组
	b := [3][2][2]int{
		{{0, 1}, {2, 3}},
		{{4, 5}, {6, 7}},
		{{8, 9}, {10, 11}},
	}
	fmt.Println(b)

	// 也可以省略各个位置的初始化,在后续代码中赋值
	c := [3][3][3]int{}
	c[2][2][1] = 5
	c[1][2][1] = 4
	fmt.Println("c = ", c)
}

func method4() {
	// 三维数组
	a := [3][2][2]int{
		{{0, 1}, {2, 3}},
		{{4, 5}, {6, 7}},
		{{8, 9}, {10, 11}},
	}

	layer1 := a[0]
	layer2 := a[0][1]
	element := a[0][1][1]
	fmt.Println(layer1)
	fmt.Println(layer2)
	fmt.Println(element)

	// 多维数组遍历时，需要使用嵌套for循环遍历
	for i, v := range a {
		fmt.Println("index = ", i, "value = ", v)
		for j, inner := range v {
			fmt.Println("inner, index = ", j, "value = ", inner)
		}
	}
}

type Custom struct {
	i int
}

var carr [5]*Custom = [5]*Custom{
	{6},
	{7},
	{8},
	{9},
	{10},
}

/*
*
数组作为参数
*/

// 值类型
// 本质：值传递，拷贝整个数组
// 特点：
//   - 参数是 [5]int，函数调用时 会复制整个数组；
//   - 函数内部修改的是副本，不会影响原数组；
//   - 如果数组很大，性能上不如传指针。
func receiveArray(param [5]int) {
	fmt.Println("in receiveArray func, before modify, param = ", param)
	param[1] = -5
	fmt.Println("in receiveArray func, after modify, param = ", param)
}

// 值类型
// 引用传递，传的是数组地址
//   - 参数是 *[5]int，传入的是数组的地址；
//   - 函数内操作的是数组本体；
//   - 修改会影响原数组。
func receiveArrayPointer(param *[5]int) {
	fmt.Println("in receiveArrayPointer func, before modify, param = ", param)
	param[1] = -5
	fmt.Println("in receiveArrayPointer func, after modify, param = ", param)
}

// 值传递数组（指针元素）
//   - [5]*Custom 是一个包含 5 个指针的数组；
//   - 整个数组是值传递，但数组中每个元素是指针；
//   - 修改指针指向的内容（即 *param[i]）是有效的，会影响原对象；
//   - 修改数组结构本身（比如替换某个指针地址）不会影响原数组。
func printFuncParamPointer(param [5]*Custom) {
	for i := range param {
		fmt.Printf("in printFuncParamPointer func, param[%d] = %p, value = %v \n", i, &param[i], *param[i])
	}
}
