package variable

import "fmt"

// 全局变量
var name string = "lixiwen"
var age int = 2
var flag bool = true
var float float64 = 3.142131

// 多变量声明
var (
	i  int  = 123
	b2 bool = false
	s2      = "hello world"
)

// 局部变量
func testVariable() {
	// 方式1，类型推导，用得最多
	a := 1
	// 方式2，完整的变量声明写法
	var b int = 2
	// 方式3，仅声明变量，但是不赋值，
	var c int
	// 多变量的定义
	var l, m, k = 1, 2, 3
	fmt.Println(a, b, c)
	fmt.Println(l, m, k)
}

// 直接在返回值中声明, 看下面我可以声明2个我的返回类型
func method2() (a int, b string) {
	return 1, "test"
}

func main() {
	testVariable()
}
