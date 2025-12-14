package operator

import "fmt"

func method() {
	a, b := 1, 2
	sum := a + b
	sub := a - b
	mul := a * b
	div := a / b
	mod := a % b

	fmt.Println(sum, sub, mul, div, mod)

}

// 当不同的算术类型混合计算时, 必须先把它们转换成同一类型才可以计算
func method2() {
	a := 10 + 0.1
	b := byte(1) + 1
	fmt.Println(a, b)

	// 先要把 b也转换为 float64才能进行计算
	sum := a + float64(b)
	fmt.Println(sum)

	// 把 a 转换为 byte类型进行计算
	var sub = byte(a) - b
	fmt.Println(sub)

	mul := a * float64(b)
	div := byte(a) / b

	fmt.Println(mul, div)

}

// 逻辑判断
func method3() {
	a := 1
	b := 5

	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
}

// 且 和 或
func method4() {
	a := true
	b := false

	fmt.Println(a && b)
	fmt.Println(a || b)
	fmt.Println(!(a && b))
}

// 位运算
func method5() {
	fmt.Println(0 & 0)
	fmt.Println(0 | 0)
	fmt.Println(0 ^ 0)

	fmt.Println(0 & 1)
	fmt.Println(0 | 1)
	fmt.Println(0 ^ 1)

	fmt.Println(1 & 1)
	fmt.Println(1 | 1)
	fmt.Println(1 ^ 1)

	fmt.Println(1 & 0)
	fmt.Println(1 | 0)
	fmt.Println(1 ^ 0)
}
