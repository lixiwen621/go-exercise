package goroutine

import (
	"fmt"
	"sync"
	"time"
)

var (
	count int
	mutex sync.Mutex
)

// Product 生产者 - 消费者 模型
func Product(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // 表示当前 goroutine 执行完成
	for i := 0; i < 10; i++ {
		fmt.Printf("[生产者] 生产数据: %d\n", i)
		ch <- i // 向通道发送数据 (发送方)
		time.Sleep(200 * time.Millisecond)
	}
	// 请注意: 由发送方关闭通道, 代表生产完成，不再发送数据
	close(ch)
}

func Consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range ch {
		fmt.Printf("[消费者] 消费数据: %d\n", val)
		time.Sleep(300 * time.Millisecond)
	}
	fmt.Println("[消费者] 通道已关闭，退出。")
}

// SyncMutex 加锁的示例方法
func SyncMutex() {
	for i := 0; i < 5; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				mutex.Lock()
				count++
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("Count", count)
}

// ProductAndConsumerTest 生产者-消费者模型 调用示例
func ProductAndConsumerTest() {
	// 1. 创建有缓冲的通道, 异步
	ch := make(chan int, 3)
	var wg sync.WaitGroup
	wg.Add(2)

	// 2. 启动生产者
	go Product(ch, &wg)
	// 3. 启动消费者
	go Consumer(ch, &wg)

	// 4. 等待所有 goroutine 完成
	wg.Wait()
	fmt.Println("所有任务完成，主程序退出")
}
