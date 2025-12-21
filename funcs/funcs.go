package main

import "fmt"

// Func0 单一返回值, 返回 hello, world
func Func0(name string) string {
	return "hello, world"
}

// Func1 多个返回值，返回 "", nil
func Func1(a, b, c int, str1 string) (string, error) {
	return "", nil
}

// Func2 带名字的返回值，返回hello <nil>
func Func2(a int, b int) (str string, err error) {
	str = "hello"
	// 带名字的返回值，可以直接 return
	return
}

// Func3 带名字的返回值，返回 hello nil
func Func3(a int, b int) (str string, err error) {
	res := "hello"
	// 虽然带名字，但是我们并没有用
	return res, nil
}

func Invoke() {
	str, err := Func2(1, 2)
	println(str, err)
	// 忽略返回值
	_, _ = Func2(1, 3)

	// 部分忽略返回值
	// str 是已经声明好了
	str, _ = Func2(3, 4)
	// str1 是新变量，需要使用 :=
	str1, _ := Func2(3, 4)
	println(str1)

	// str2 是新变量，需要使用 :=
	str2, err := Func2(3, 4)
	println(str2)
}

// Recursive 递归
// 递归使用不当, 就有可能出现 stack overflow
func Recursive() {
	Recursive()
}

// Func4 演示了 Go 中函数作为一等公民（first-class citizen）的用法：函数可以像普通值一样被赋值、传递和存储。
func Func4() {
	// 直接调用
	result, err := Func3(1, 2)
	// 赋值后调用（效果相同）
	myFunc3 := Func3
	result2, err2 := myFunc3(1, 2)
	fmt.Println(result, err)
	fmt.Println(result2, err2)
}

// 也是一等公民的特性
// 函数变量
func Func41() {
	var f func(int, int) int
	f = func(x, y int) int {
		return x + y
	}
	fmt.Println(f(2, 3)) // 输出 5

	// 实际应用 这种模式常用于：
	// 回调函数
	// 策略模式（根据不同条件选择不同算法）
	//函数式编程（map、filter、reduce 等操作）
	// 例如，可以动态选择不同的计算方式
	var calculator func(int, int) int
	operation := "add" // 定义 operation 变量

	if operation == "add" {
		calculator = func(x, y int) int { return x + y }
	} else if operation == "multiply" {
		calculator = func(x, y int) int { return x * y }
	}

	result := calculator(2, 3)
	fmt.Println(result) // 使用 result 变量
}

// 这是一个高阶函数（Higher-Order Function）示例：函数作为参数传递
func operate(x, y int, op func(int, int) int) int {
	return op(x, y)
}

// 也是一等公民的特性
// 调用上面的 operate 方法示例
// 函数作为参数
func Fun42() {
	// 示例 1: 传入加法函数
	add := func(x, y int) int { return x + y }
	result1 := operate(3, 4, add) // 结果: 7
	fmt.Println(result1)
	// 示例 2: 传入乘法函数
	multiply := func(x, y int) int { return x * y }
	result2 := operate(3, 4, multiply) // 结果: 12
	fmt.Println(result2)
	// 示例 3: 直接传入匿名函数
	result3 := operate(10, 5, func(x, y int) int { return x - y }) // 结果: 5
	fmt.Println(result3)
	// 示例 4: 传入除法函数
	result4 := operate(10, 2, func(x, y int) int { return x / y }) // 结果: 5
	fmt.Println(result4)
}

// 函数作为返回值
func getAdder(base int) func(int) int {
	return func(v int) int {
		return base + v
	}
}

func Func43() {
	adder := getAdder(10)
	fmt.Println(adder(1)) // 输出 11
	fmt.Println(adder(2)) // 输出 13
	fmt.Println(adder(3)) // 输出 16
}

func Func5() {
	fn := func(name string) string {
		return "hello, " + name
	}

	str := fn("大明")
	fmt.Println(str)
}

// Func6 的返回值是一个方法，
func Func6() func(name string) string {
	return func(name string) string {
		return "hello," + name
	}
}

func Func6Invoke() {
	// sayHello 的签名就是 func(name string) string
	sayHello := Func6()
	sayHello("大明")
}

// Func7 演示匿名方法
func Func7() {
	hello := func() string {
		return "hello, world"
	}()
	println(hello)
}

// 闭包
func Closure(name string) func() string {

	// 返回的这个函数，就是一个闭包。
	// 它引用到了 Closure 这个方法的入参
	return func() string {
		return "hello, " + name
	}
}

// YourName 不定参数例子
// 一个人可能有很多个别名，也可能一个都没有
func YourName(name string, alias ...string) {
	if len(alias) > 0 {
		println(alias[0])
	}
}

func YourNameInvoke() {
	YourName("Deng Ming")
	YourName("Deng Ming", "Da Ming")
	YourName("Deng Ming", "Da Ming", "Zhong Ming")
}

func Defer() {
	defer func() {
		println("第一个 defer")
	}()

	defer func() {
		println("第二个 defer")
	}()
}

func DeferNum() {
	for i := 0; i < 100; i++ {
		defer func() {
			println("hello, world")
		}()
	}
}

func DeferOpenCode() {
	defer func() {
		println("hello, world")
	}()

	// 假如说这个是你的业务代码
	YourBusiness()

	// 开放编码就类似于编译器帮你把 defer 的内容放过来了这里
	//func() {
	//	println("hello, world")
	//}()
}

func YourBusiness() {
	// 假如说这是你的业务代码
}

func DeferConditional(input bool) {
	if input {
		defer func() {
			println("hello, world")
		}()
	}
}

func DeferClosure() {
	i := 0
	defer func() {
		println(i)
	}()
	i = 1
}

func DeferClosureV1() {
	i := 0
	defer func(val int) {
		println(val)
	}(i)
	i = 1
}

func DeferClosureLoopV1() {
	for i := 0; i < 10; i++ {
		defer func() {
			println(i)
		}()
	}
}

func DeferClosureLoopV2() {
	for i := 0; i < 10; i++ {
		defer func(val int) {
			println(val)
		}(i)
	}
}

func DeferClosureLoopV3() {
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			println(j)
		}()
	}
}

func DeferReturn() int {
	a := 0
	defer func() {
		a = 1
	}()
	return a
}

func DeferReturnV1() (a int) {
	a = 0
	defer func() {
		a = 1
	}()
	return a
}

func DeferReturnV2() *MyStruct {
	a := &MyStruct{
		name: "Jerry",
	}
	defer func() {
		a.name = "Tom"
	}()

	return a
}

type MyStruct struct {
	name string
}
