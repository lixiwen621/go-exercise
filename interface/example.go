package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

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

// golang中的接口可以隐式的实现
func method1(p Speaker) {
	var s Speaker = Cat{}
	fmt.Println(s.Speak())
}

// 接口的类型断言和类型判断
func method2() {
	var s Speaker = Cat{}
	c, ok := s.(Cat)
	if ok {
		fmt.Println(c.Speak())
	}

	switch v := s.(type) {
	case Cat:
		fmt.Println("是猫")
		fmt.Println(v.Speak())
	case Speaker:
		fmt.Println("是接口")
		fmt.Println(v.Speak())
	}
}

// 空接口 万能类型
func method3() {
	// 定义一个包含多种类型的 Map, 相当于Java中的 Map<String, Object>
	data := make(map[string]interface{})
	data["name"] = "Tom"
	data["age"] = 18
	data["married"] = false
	data["hobbies"] = []string{"reading", "cycling"}

	for key, value := range data {
		fmt.Printf("Key: %s, Value: %v, Type: %T\n", key, value, value)
	}
}

// 空接口 使用 interface{} 进行解码
func method4() {
	jsonStr := `{
        "name": "Bob",
        "age": 25,
        "skills": ["Go", "Python", "Docker"]
    }`

	var result map[string]interface{}
	// 解码 JSON
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// 访问解码后的数据
	fmt.Println("Name:", result["name"])
	fmt.Println("Age:", result["age"])

	// 类型断言提取 skills 切片
	// 请注意下面的语法 if val, ok := expr; ok { ... } -》等价于 val, ok := expr  if ok { ... }
	if skills, ok := result["skills"].([]interface{}); ok {
		fmt.Println("Skills:")
		for _, skill := range skills {
			fmt.Println("-", skill)
		}
	}
}

// 接口组合和继承
// 接口可以嵌套组合，实现类似继承的效果
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

// readSink 只接受「读」侧能力：参数类型为 Reader，不关心对方是否还能 Write。
func readSink(r Reader) {
	buf := make([]byte, 64)
	n, err := r.Read(buf)
	fmt.Printf("[readSink] 读取 %d 字节: %q err=%v\n", n, buf[:n], err)
}

// writeSink 只接受「写」侧能力。
func writeSink(w Writer) {
	n, err := w.Write([]byte(" -> writeSink 写入\n"))
	fmt.Printf("[writeSink] 写入 %d 字节 err=%v\n", n, err)
}

// ReadWriterDemo 演示接口组合：ReadWriter = Reader ∪ Writer。
//
// *bytes.Buffer 自带的 Read/Write 与上面 Reader、Writer 的方法签名一致，
// 因此同时满足 Reader、Writer，自然也满足把它们嵌在一起的 ReadWriter。
func ReadWriterDemo() {
	buf := bytes.NewBuffer(nil)
	var rw ReadWriter = buf

	if _, err := rw.Write([]byte("Hello")); err != nil {
		fmt.Println("Write:", err)
		return
	}
	if _, err := rw.Write([]byte("，组合接口")); err != nil {
		fmt.Println("Write:", err)
		return
	}

	readBuf := make([]byte, 64)
	n, err := rw.Read(readBuf)
	fmt.Printf("[ReadWriter] 读取: %q (%d 字节) err=%v\n", readBuf[:n], n, err)

	fmt.Println("--- 收窄：把 ReadWriter 传给只要 Reader 或只要 Writer 的函数 ---")
	buf2 := bytes.NewBuffer([]byte("abc"))
	var rw2 ReadWriter = buf2
	writeSink(rw2)
	readSink(rw2)
}

// 使用接口的相关示例
// PaymentMethod 接口定义了支付方法的基本操作
type PayMethod interface {
	Account //账户
	Pay(amount int) bool
}

type Account interface {
	GetBalance() int // 获取余额
}

// CreditCard 结构体实现 PaymentMethod 接口
// 信用卡
type CreditCard struct {
	balance int //余额
	limit   int //限制
}

func (c *CreditCard) Pay(amount int) bool {
	if c.balance+amount <= c.limit {
		c.balance += amount
		fmt.Printf("信用卡支付成功: %d\n", amount)
		return true
	}
	fmt.Println("信用卡支付失败: 超出额度")
	return false
}

// 获取信用卡余额
func (c *CreditCard) GetBalance() int {
	return c.balance
}

// DebitCard 结构体实现 PaymentMethod 接口
// 借记卡
type DebitCard struct {
	balance int
}

func (d *DebitCard) Pay(amount int) bool {
	if d.balance >= amount {
		d.balance -= amount
		fmt.Printf("借记卡支付成功: %d\n", amount)
		return true
	}
	fmt.Println("借记卡支付失败: 余额不足")
	return false
}

// 获取借记卡余额
func (d *DebitCard) GetBalance() int {
	return d.balance
}

// 使用 PaymentMethod 接口的函数
func purchaseItem(p PayMethod, price int) {
	if p.Pay(price) {
		fmt.Printf("购买成功，剩余余额: %d\n", p.GetBalance())
	} else {
		fmt.Println("购买失败")
	}
}

func anyParam(param interface{}) {
	fmt.Println("param: ", param)
}

func method5() {
	creditCard := &CreditCard{balance: 0, limit: 1000}
	debitCard := &DebitCard{balance: 500}

	fmt.Println("使用信用卡购买:")
	purchaseItem(creditCard, 800)

	fmt.Println("\n使用借记卡购买:")
	purchaseItem(debitCard, 300)

	fmt.Println("\n再次使用借记卡购买:")
	purchaseItem(debitCard, 300)

	fmt.Print("信用卡的参数 ")
	anyParam(creditCard)
	fmt.Print("借记卡的参数 ")
	anyParam(debitCard)

}
