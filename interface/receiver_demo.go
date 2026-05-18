package main

import "fmt"

// Dog 仅实现指针接收者的 Speak，用于对比 example.go 里 Cat 的值接收者。
type Dog struct{}

func (d *Dog) Speak() string {
	return "Woof"
}

// ReceiverWithInterfaceDemo 演示：值接收者 vs 指针接收者对接口赋值的影响。
// - func (T) M()：T{} 与 *T{} 一般都可赋给「需要 M 的接口」
// - func (*T) M()：通常只有 *T{} 可以，T{} 不行（T 的方法集里没有 M）
func ReceiverWithInterfaceDemo() {
	fmt.Println("=== Cat：值接收者 (c Cat) Speak — Cat{} 与 &Cat{} 均可赋给 Speaker ===")
	var sCat1 Speaker = Cat{}
	var sCat2 Speaker = &Cat{}
	fmt.Printf("  Speaker = Cat{}   -> %q\n", sCat1.Speak())
	fmt.Printf("  Speaker = &Cat{}  -> %q\n", sCat2.Speak())

	fmt.Println("=== Dog：指针接收者 (d *Dog) Speak — 仅 &Dog{} 可赋给 Speaker ===")
	var sDog Speaker = &Dog{}
	fmt.Printf("  Speaker = &Dog{}  -> %q\n", sDog.Speak())

	// 下面一行若取消注释会编译失败：Dog 的方法集里没有 Speak（方法只定义在 *Dog 上）
	// var sBad Speaker = Dog{}
	// _ = sBad

	fmt.Println("=== example2：SayHello 只定义在 *Student / *Teacher 上 ===")
	whoSayHello(&Student{})
	whoSayHello(&Teacher{})

	// 下面两行若取消注释会编译失败：Student / Teacher 值类型的方法集里没有 SayHello
	// whoSayHello(Student{})
	// whoSayHello(Teacher{})
}
