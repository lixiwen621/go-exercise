package constants

const age2 = 12

// 定义枚举
// 枚举的本质就是一系列常量
const (
	Male   = "Male"
	Female = "Female"
)

// 自定义命名类型（与 string 不同，需显式转换才能与 string 互换）
type Gender string

const (
	Male2   Gender = "male"
	Female2 Gender = "female"
)

// 上面的枚举作为参数的时候，会使用Gender作为参数类型
func method(gender Gender) {

}

// 给Gender类型，自定义 string() 方法
// 请注意 *Gender 是方法的接收者类型，后续在函数和方法的章节会介绍
// *Gender类型是 指针
func (g *Gender) String() string {
	if g == nil {
		return ""
	}
	switch *g {
	case Male2:
		return string(Male2)
	case Female2:
		return string(Female2)
	default:
		return "Unknown"
	}
}

func (g *Gender) IsMale() bool {
	if g == nil {
		return false
	}
	return *g == Male2
}

// iota关键字自动为常量赋值
type ConnState int

const (
	StateNew ConnState = iota
	StateActive
	StateIdle
	StateClosed
)

type Month int

const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)
