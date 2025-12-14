package goroutine

import (
	"fmt"
	"sync"
	"time"
)

// 简单示例
func sayHello(name string) {
	fmt.Println("Hello,", name)
}

func method1() {
	// 在func 前面加上 go, 它就会并发的执行
	go sayHello("Alice")
	go sayHello("Bob")

	// 等一等，不然 main 结束后 goroutine 会直接被结束
	time.Sleep(time.Second)
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func method2() {
	// 匿名函数
	go func() {
		fmt.Println("run goroutine in closure")
	}()
	// 带参数的匿名函数
	go func(string) {
	}("gorouine: closure params")
	go say("in goroutine: world")
	say("hello")
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

func method3() {
	counter := UnSafeCounter{}

	// 100个goroutine 同时增加技术
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				counter.Increment()
			}
		}()
	}

	// 等待一段时间确保所有goroutine完成
	time.Sleep(time.Second)

	// 输出最终计数
	fmt.Printf("Final count: %d\n", counter.GetCount())
}

// channel 通道的用法
func method4() {
	ch := make(chan string)

	// 发送方
	go func() {
		ch <- "Hello from goroutine"
	}()

	// 接收方(会阻塞直到收到数据)
	msg := <-ch
	fmt.Println(msg)
}

// 无缓冲通道(同步)
func method5() {
	ch := make(chan int)
	go func() {
		ch <- 1 // 阻塞直到有接受
	}()
	fmt.Println(<-ch) // 接受后，发送方解除阻塞
}

// 有缓存通道(异步)
func method6() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
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

func method7() {
	ch := make(chan int, 3)
	go producer(ch)
	// 主 goroutine 作为消费者
	consumer(ch)
}

// select 用于同时监听多个 channel，哪个有数据就先处理哪个
func method8() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() { ch1 <- "Hello from ch1" }()
	go func() { ch2 <- "Hello from ch2" }()

	timeout := time.After(2 * time.Second)
	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		case <-timeout:
			fmt.Println("操作超时")
			return
		}
	}
}

// 稍微复杂点的示例
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

func method9() {
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
