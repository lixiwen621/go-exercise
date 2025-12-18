package _map

import (
	"fmt"
	"sync"
	"time"
)

// 声明 和 初始化
func method1() {
	// key为 int类型, value为string类型
	// 声明 map 函数
	// 请注意, 不能直接写未初始化的 map会panic
	var m1 map[int]string
	fmt.Println(m1)

	// 使用make函数初始化 map
	m2 := make(map[string]string)
	fmt.Println(m2)

	// 初始化的时候定义 容量为10
	m3 := make(map[string]string, 10)
	fmt.Println("m3 length:", len(m3))
	fmt.Println(m3)

	// 声明和初始化 m4的 map函数
	m4 := map[string]string{}
	fmt.Println("m4 length:", len(m4))
	fmt.Println("m4 =", m4)
	m4["lixiwen"] = "tongguo.li"
	fmt.Println("m4 length:", len(m4))

	// 初始化 和添加 元素
	m5 := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	fmt.Println("m5 length:", len(m5))
	fmt.Println("m5 =", m5)

}

// 内置的 map 删除函数
func method2() {
	m := make(map[string]int, 100)
	fmt.Println(m)

	m["dog"] = 1
	m["cat"] = 2
	m["dog"] = 3 // 覆盖旧值

	// 使用内置函数 delete() 删除 map 集合中指定的 key
	delete(m, "cat")
	fmt.Println(m)
}

// 检查map中的 键是否存在
func method3() {
	m := make(map[string]int, 100)
	m["cat"] = 1
	v, ok := m["cat"]
	if ok {
		fmt.Println("Key exists with value:", v)
	} else {
		fmt.Println("Key not found")
	}
}

func method4() {
	m := make(map[string]int, 10)

	m["1"] = int(1)
	m["2"] = int(2)
	m["3"] = int(3)
	m["4"] = int(4)
	m["5"] = int(5)
	m["6"] = int(6)

	// 获取元素
	value1 := m["1"]
	fmt.Println("m[\"1\"] =", value1)

	// 存在的话 exist会返回为 true
	value1, exist := m["1"]
	fmt.Println("m[\"1\"] =", value1, ", exist =", exist)

	valueUnexist, exist := m["10"]
	fmt.Println("m[\"10\"] =", valueUnexist, ", exist =", exist)

	// 修改值
	fmt.Println("before modify, m[\"2\"] =", m["2"])
	m["2"] = 20
	fmt.Println("after modify, m[\"2\"] =", m["2"])

	// 获取map的长度
	fmt.Println("before add, len(m) =", len(m))
	m["10"] = 10
	fmt.Println("after add, len(m) =", len(m))

	// 遍历map集合main
	for key, value := range m {
		fmt.Println("iterate map, m[", key, "] =", value)
	}

	// 使用内置函数删除指定的key
	_, exist_10 := m["10"]
	fmt.Println("before delete, exist 10: ", exist_10)
	delete(m, "10")
	_, exist_10 = m["10"]
	fmt.Println("after delete, exist 10: ", exist_10)

	// 在遍历时，删除map中的key
	for key := range m {
		fmt.Println("iterate map, will delete key:", key)
		delete(m, key)
	}
	fmt.Println("m = ", m)
}

// map 集合作为参数的时候
// 对 map 集合的修改，都会影响到原始实参
func method5() {
	m := make(map[string]int)
	m["a"] = 1
	receiveMap(m)
	fmt.Println("m =", m)

}

func receiveMap(param map[string]int) {
	fmt.Println("before modify, in receiveMap funcs: param[\"a\"] = ", param["a"])
	param["a"] = 2
	param["b"] = 3
}

// 并发时使用 map
// go中的map不是线程安全的 并发读写 map 是不安全的
// 请注意 method6()里的代码会报异常  fatal error: concurrent map read and map write
// 想要正确的执行下面的方法 必须加锁, 因map不是线程安全的
func method6() {
	m := make(map[string]int)
	// 启动第一个 goroutine，不断对 map 中 "a" 对应的值进行递增操作
	// 这其实是一个无限循环中的 写操作
	go func() {
		for {
			m["a"]++
		}
	}()

	// 这里是一个无限循环的 读写操作 fmt.Println(m["a"])是读
	go func() {
		for {
			m["a"]++
			fmt.Println(m["a"])
		}
	}()
	// 主协程阻塞 5 秒后超时，打印 "timeout, stopping"，然后程序退出
	// 在这段时间内，两个 goroutine 会并发运行
	select {
	case <-time.After(time.Second * 5):
		fmt.Println("timeout, stopping")
	}
}

// method6()方法中的加锁操作
func method7() {
	m := make(map[string]int)
	var lock sync.Mutex
	go func() {
		for {
			lock.Lock()
			m["a"]++
			lock.Unlock()
		}
	}()

	go func() {
		for {
			lock.Lock()
			m["a"]++
			fmt.Println(m["a"])
			lock.Unlock()
		}
	}()

	select {
	case <-time.After(time.Second * 5):
		fmt.Println("timeout, stopping")
	}
}
