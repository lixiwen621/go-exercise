package constants

import "fmt"

// 方式1:
const a int = 1

// 方式2:
const name = "tom"
const age = 30

// 方式3:
const c, d = 2, "hello"

// 方式4:
const e, f bool = false, true

// 方式5:
const (
	h    byte = 3
	i         = "hello worlds"
	j, k      = 4, 'v'
	l, m      = 5, false
)

// 方式6:
const (
	fullName string = "lixiwen"
	grade    int    = 12
)

func main() {
	fmt.Println(name, age)
}
