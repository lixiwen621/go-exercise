package control

import "fmt"

// 简单的 if-else 流程语句判断
func method() {
	age := 18
	if age < 18 {
		fmt.Println("未成年")
	} else if age == 18 {
		fmt.Println("刚成年")
	} else {
		fmt.Println("成年人")
	}

}

// 可初始化值的 if-else 流程语句判断
func method2() {
	var a int = 10
	if b := 1; a > 10 {
		fmt.Println(b)
		fmt.Println("a > 10")
	} else if c := 3; b > 1 {
		b = 3
		fmt.Println("b > 1")
	} else {
		fmt.Println("其他")
		if c == 3 {
			fmt.Println("c == 3")
		}
		fmt.Println(b)
		fmt.Println(c)
	}
}

// switch 语句示例
func method3() {
	a := "test string"
	switch a {
	case "test":
		fmt.Println("a= ", a)
	case "string":
		fmt.Println("a= ", a)
	case "s", "test string":
		fmt.Println("a= ", a)
	}
}

// switch中也可以先 初始化变量
func method4() {
	// 变量b仅在当前switch代码块内有效
	switch b := 5; b {
	case 1:
		fmt.Println("b = 1")
	case 2:
		fmt.Println("b = 2")
	case 3, 4:
		fmt.Println("b = 3 or 4")
	case 5:
		fmt.Println("b = 5")
	default:
		fmt.Println("b = ", b)
	}
}

func method5() {
	// 不指定判断变量，直接在case中添加判定条件
	a := "test string"
	b := 5
	switch {
	case a == "t":
		fmt.Println("a = t")
	case b == 3:
		fmt.Println("b = 5")
	case b == 5, a == "test string":
		fmt.Println("a = test string; or b = 5")
	default:
		fmt.Println("default case")
	}
}

type CustomType struct {
}

func method6() {
	var d interface{}
	// var e byte = 1
	d = 1
	switch t := d.(type) {
	case byte:
		fmt.Println("d is byte type, ", t)
	case *byte:
		fmt.Println("d is byte point type, ", t)
	case *int:
		fmt.Println("d is int type, ", t)
	case *string:
		fmt.Println("d is string type, ", t)
	case *CustomType:
		fmt.Println("d is CustomType pointer type, ", t)
	case CustomType:
		fmt.Println("d is CustomType type, ", t)
	default:
		fmt.Println("d is unknown type, ", t)
	}
}
