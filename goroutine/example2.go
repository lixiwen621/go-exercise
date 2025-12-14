package goroutine

import (
	"fmt"
	"time"
)

func printNumber() {
	for i := 1; i <= 5; i++ {
		fmt.Println("数字:", i)
		time.Sleep(200 * time.Millisecond)
	}
}

func worker(ch chan int) {
	// 接受数据
	fmt.Println("收到:", <-ch)
}

// Send 发送数据的方法
func Send() {
	// 无缓冲通道
	ch := make(chan int)
	go worker(ch)
	ch <- 42 // 发送数据到通道

	// 缓冲通道, 缓冲通道为3
	ch2 := make(chan int, 3)

	go worker(ch2)
	ch2 <- 12
	ch2 <- 13
	ch2 <- 14
	ch2 <- 11
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
}

// Close 关闭通道
func Close() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// 关闭通道后 不能在发送数据, 但是可以接受数据
	close(ch)
	for val := range ch {
		fmt.Println(val)
	}
}

// SelectTest 多路复用的示例方法
func SelectTest() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	// 匿名函数 前面加 go就是异步执行
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "来自通道1"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "来自通道2"
	}()

	// 使用 select 等待多个通道
	select {
	case msg1 := <-ch1:
		fmt.Println("收到:", msg1)
	case msg2 := <-ch2:
		fmt.Println("收到:", msg2)
	case <-time.After(3 * time.Second):
		fmt.Println("超时: 3秒内没有收到消息")
	}
}

func Main() {
	go printNumber() // 启动一个新的goroutine
	fmt.Println("主函数运行中...")

	time.Sleep(1 * time.Second) // 主程序等待一段时间, 防止主程序提前退出
	fmt.Println("程序结束")

}
