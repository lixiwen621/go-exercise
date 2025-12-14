package typeconversion

import (
	"fmt"
	"strconv"
)

func method1() {
	var i int32 = 17
	var b byte = 5
	var f float32

	// 数字类型可以直接强转
	f = float32(i) / float32(b)
	fmt.Printf("%f\n", f)

	// 当 int32 转换为 byte时, 高位被直接舍弃
	var i2 int32 = 256
	var b2 byte = byte(i2)
	fmt.Printf("b2 的值为: %d\n", b2)
}

func method11() {
	var i int32 = 17
	var b byte = 5
	var c int32 = 18
	// i += b // 会报编译出错 因为不是相同的类型
	i += int32(b)
	i += c

}

// 字符串类型转换
func method2() {
	str := "hello, 123, 你好"
	var bytes []byte = []byte(str)
	var runes []rune = []rune(str)
	fmt.Printf("bytes 的值为: %v \n", bytes)
	fmt.Printf("runes 的值为: %v \n", runes)

	str2 := string(bytes)
	str3 := string(runes)
	fmt.Printf("str2 的值为: %v \n", str2)
	fmt.Printf("str3 的值为: %v \n", str3)
}

// golang中经常用的 strconv库来进行 字符串与数字的转换
func method3() {
	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	fmt.Printf("字符串转换为int: %d \n", num)
	str1 := strconv.Itoa(num)
	fmt.Printf("int转换为字符串: %s \n", str1)
	// 方法中 base为10 就是转换为10进制
	// Parse一类的方法就是 把字符串转换为 对应类型
	ui64, err := strconv.ParseUint(str1, 10, 64)
	fmt.Printf("字符串转换为uint64: %d \n", ui64)
	//方法中 base为2 就是转换为 2进制
	// Format 一类的方法就是 把各个类型转换为 字符串
	str2 := strconv.FormatUint(ui64, 2)
	fmt.Printf("uint64转换为字符串: %s \n", str2)

}

// 接口类型的转换
func method4() {
	var i interface{} = 3
	a, ok := i.(int)
	if ok {
		fmt.Printf("'%d' is a int \n", a)
	} else {
		fmt.Println("conversion failed")
	}
}

// 通过 switch 来判断
func method5() {
	var i interface{} = 3
	switch v := i.(type) {
	case int:
		fmt.Printf("i is a int \n", v)
	case string:
		fmt.Printf("i is a string \n", v)
	default:
		fmt.Printf("i is unknown type \n", v)
	}
}

// 自定义接口的 转换
// 下面3个(包括接口 结构体 函数的实现) 是一种隐式接口的实现
type Supplier interface {
	Get() string
}

type DigitSupplier struct {
	value int
}

func (i *DigitSupplier) Get() string {
	return fmt.Sprintf("%d", i.value)
}

func method6() {
	// 结构体初始化的前面加 &就变成了 取地址操作
	var a Supplier = &DigitSupplier{value: 1}
	fmt.Println(a.Get())
	// 尝试将接口变量 a 的底层类型转换为 *DigitSupplier 类型（指针类型）
	b, ok := (a).(*DigitSupplier)
	fmt.Println(b, ok)
}
