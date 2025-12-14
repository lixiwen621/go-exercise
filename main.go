package main

import (
	"fmt"
	_map "go-exercise/map"
	"sync"
	"time"
)

const mainName string = "main"

var mainVar string = getMainVar()

// 初始化方法
func init() {
	fmt.Println("main init method invoked")
}

func main() {
	fmt.Println("main method invoked!")
	_map.Test()
}

func getMainVar() string {
	fmt.Println("main.getMainVar method invoked!")
	return mainName
}

// 定义一个接口
type Speaker interface {
	Speak() string
}

// 隐式实现接口
type Cat struct {
}

func (c Cat) Speak() string {
	return "Tom"
}

// 增加安全的计数器
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

// 增加计数
func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

// 获取当前计数
func (c *SafeCounter) GetCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

// 不安全的计数器
type UnSafeCounter struct {
	count int
}

// 增加计数
func (c *UnSafeCounter) Increment() {
	c.count++
}

// 获取计数
func (c *UnSafeCounter) GetCount() int {
	return c.count
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// 生产者-消费者模式
func producer(ch chan int) {
	for i := 0; i < 5; i++ {
		fmt.Println("生产:", i)
		ch <- i
		time.Sleep(time.Millisecond * 200)
	}
	close(ch) // 生产完毕, 关闭通道
}

func consumer(ch chan int) {
	for val := range ch {
		fmt.Println("消费:", val)
	}
}

// 只接收channel的函数
func receiveOnly(ch <-chan int) {
	for v := range ch {
		fmt.Printf("接收到: %d\n", v)
	}
}

// 只发送channel的函数
func sendOnly(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Printf("发送: %d\n", i)
	}
	close(ch)
}

func method2() {
	// 创建一个带缓冲的channel
	ch := make(chan int, 3)

	// 启动发送goroutine
	go sendOnly(ch)

	// 启动接收goroutine
	go receiveOnly(ch)

	// 使用select进行多路复用
	timeout := time.After(2 * time.Second)
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Println("Channel已关闭")
				return
			}
			fmt.Printf("主goroutine接收到: %d\n", v)
		case <-timeout:
			fmt.Println("操作超时")
			return
		default:
			fmt.Println("没有数据，等待中...")
			time.Sleep(500 * time.Millisecond)
		}
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

func initInt() {
	var a1 int = 1
	var b1 int = 2

	// 十六进制
	var a uint8 = 0xF
	var b uint8 = 0xf

	// 八进制
	var c uint8 = 017
	var d uint8 = 0o17
	var e uint8 = 0o17

	// 二进制
	var f uint8 = 0b1111
	var g uint8 = 0b1111

	// 十进制
	var h uint8 = 15
	fmt.Println(a1, b1)
	fmt.Println(a == b)
	fmt.Println(c == d)
	fmt.Println(d == e)
	fmt.Println(f == g)
	fmt.Println(g == h)
}

func exampleFloat() {
	// 创建的32位
	var float1 float32 = 10
	// 创建的64位
	float2 := 10.0
	fmt.Println(float1)
	fmt.Println(float2)
}

func exampleComplex() {
	// 创建的complex128 位
	var x complex128 = complex(3, 4)
	// 默认创建 128位
	c1 := complex(4, 5)
	x = c1
	fmt.Println("复数:", x)
	fmt.Println("实部:", real(x))
	fmt.Println("虚部:", imag(x))

	y := 1 + 2i
	sum := x + y
	product := x * y
	divide := x / y

	fmt.Println("和: ", sum)
	fmt.Println("积:", product)
	fmt.Println("除为:", divide)
}

func exampleByte() {
	// 字符串转换为 byte[] 数组
	// byte[] 数组 也是Unicode 码点值
	var s string = "Hello World"
	var bytes []byte = []byte(s)
	fmt.Println("convert \"Hello, World!\" to bytes: ", bytes)
	// byte[] 数组 转换为字符串
	var newString string = string(bytes)
	fmt.Println(newString)
}

// 字符类型
func exampleRune() {
	// rune值 也是表示一个字符的 Unicode值
	var r1 rune = 'a'
	var r2 rune = '世'
	// 打印出来的是 Unicode 编码值
	fmt.Println(r1, r2)

	var s string = "abc, 你好, 世界!"
	var runes []rune = []rune(s)
	fmt.Println(runes)
}
