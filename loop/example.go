package loop

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

// 声明 for 循环
func method() {
	// 方式1
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 方式2
	b := 1
	for b < 10 {
		fmt.Println(b)
		b++
	}

	// 方式3，无限循环
	// 创建 context, 2秒后 超时
	// ctx.Done() 会关闭（channel 关闭），用于通知取消
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))
	// 控制变量
	var started bool
	// 原子性的 bool值
	var stopped atomic.Bool
	for {
		if !started {
			started = true
			// 启动子 goroutine
			// goroutine 会一直监听 ctx.Done()
			go func() {
				for {
					select {
					// 超过2秒, 就会执行 ctx.Done()
					case <-ctx.Done():
						fmt.Println("ctx done")
						stopped.Store(true)
						return
					}
				}
			}()
		}
		fmt.Println("main")
		if stopped.Load() {
			break
		}
	}
}

// 遍历数组
func method2() {
	var a [10]string
	a[0] = "hello"
	// 打印当前下标
	for i := range a {
		fmt.Println(i)
	}

	// 可以打印出当前下标, 和下标对应的value值
	for i, e := range a {
		fmt.Println("a[", i, "] = ", e)
	}
}

func method3() {
	// 遍历切片
	// string[] 数组切片
	s := make([]string, 10)
	s[0] = "Hello"
	// 声明一个 取下标
	for i := range s {
		fmt.Println("当前下标：", i)
	}
	// 声明两个 取下标和对应的value值
	for i, e := range s {
		fmt.Println("s[", i, "] = ", e)
	}

	// map 切片
	m := make(map[string]string)
	m["b"] = "Hello, b"
	m["a"] = "Hello, a"
	m["c"] = "Hello, c"
	for i := range m {
		fmt.Println("当前key：", i)
	}
	for k, v := range m {
		fmt.Println("m[", k, "] = ", v)
	}
}

// break 中断
func method4() {
	// 中断for循环
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println("第", i, "次循环")
	}

	// 中断switch
	// 中断的是 case里的逻辑
	switch i := 1; i {
	case 1:
		fmt.Println("进入case 1")
		if i == 1 {
			break
		}
		// 中断后就不会执行 下面的这里的句子了
		fmt.Println("i等于1")
	case 2:
		fmt.Println("i等于2")
	default:
		fmt.Println("default case")
	}

	// 中断select
	select {
	case <-time.After(time.Second * 2):
		fmt.Println("过了2秒")
	case <-time.After(time.Second):
		fmt.Println("进过了1秒")
		if true {
			break
		}
		fmt.Println("break 之后")
	}

	// 不使用标记
	for i := 1; i <= 3; i++ {
		fmt.Printf("不使用标记,外部循环, i = %d\n", i)
		for j := 5; j <= 10; j++ {
			fmt.Printf("不使用标记,内部循环 j = %d\n", j)
			break
		}
	}

	// 使用标记
outter:
	for i := 1; i <= 3; i++ {
		fmt.Printf("使用标记,外部循环, i = %d\n", i)
		for j := 5; j <= 10; j++ {
			fmt.Printf("使用标记,内部循环 j = %d\n", j)
			break outter
		}
	}
}

func method5() {
	gotoPreset := false
preset:
	a := 5

process:
	if a > 0 {
		a--
		fmt.Println("当前a的值为：", a)
		goto process
	} else if a <= 0 {
		// elseProcess:
		if !gotoPreset {
			gotoPreset = true
			goto preset
		} else {
			goto post
		}
	}

post:
	fmt.Println("main将结束，当前a的值为：", a)
}
