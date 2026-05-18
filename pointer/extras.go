package main

import "fmt"

// demoNilPointer：指针零值为 nil，解引用 nil 会 panic。
func demoNilPointer() {
	var p *int
	fmt.Println("未赋值指针为 nil:", p == nil)
	if p != nil {
		fmt.Println(*p)
	}
	// fmt.Println(*p) // 取消注释运行会 panic: nil pointer dereference
}

// demoNew：new(T) 返回 *T，指向一块已零初始化的 T。
func demoNew() {
	p := new(int)
	fmt.Println("new(int) 初始 *p =", *p)
	*p = 42
	fmt.Println("赋值后 *p =", *p)
}

type Person struct {
	Name string
	Age  int
}

// demoStructPointer：对 *Person 访问字段时，可写 p.Name，等价 (*p).Name。
func demoStructPointer() {
	p := &Person{Name: "Tom", Age: 18}
	p.Name = "Jerry"
	(*p).Age = 20
	fmt.Printf("结构体指针: %+v\n", *p)
}

// demoPointerCompare：== 比较的是地址是否相同，不是比较指向的值。
func demoPointerCompare() {
	a, b := 1, 1
	pa, pb, pc := &a, &b, &a
	fmt.Println("pa 与 pc 同指向 a:", pa == pc)
	fmt.Println("pa 与 pb 不同变量（即使值都为 1）:", pa == pb)
}

// demoMapAndAddress：map 的值不可寻址，不能 &m[key]；取临时变量地址只会改副本。
func demoMapAndAddress() {
	m := map[string]int{"x": 1}
	v := m["x"]
	pv := &v
	*pv = 999
	fmt.Println("改 *pv 后 map['x'] 仍是:", m["x"])
	m["x"] = 2
	fmt.Println("需要改 map 应直接赋值 m[\"x\"]:", m["x"])
}

// demoSliceElementAddress：切片元素在内存中连续，可取 &s[i] 并在适当时机修改原切片。
func demoSliceElementAddress() {
	s := []int{10, 20}
	p0 := &s[0]
	*p0 = 99
	fmt.Println("通过 &s[0] 修改后:", s)
}

// RunPointerExtras 汇总运行上述补充示例（与 defined_pointer.go 基础示例配套）。
func RunPointerExtras() {
	fmt.Println("=== nil 指针 ===")
	demoNilPointer()
	fmt.Println("=== new(T) ===")
	demoNew()
	fmt.Println("=== 结构体指针 ===")
	demoStructPointer()
	fmt.Println("=== 指针比较（地址）===")
	demoPointerCompare()
	fmt.Println("=== map 与取址 ===")
	demoMapAndAddress()
	fmt.Println("=== 切片元素取址 ===")
	demoSliceElementAddress()
}
