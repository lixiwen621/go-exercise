package pkg1

import (
	"fmt"
	_ "github.com/lixiwen621/go-exercise/pkg2"
)

const PkgName string = "pkg1"

var PkgNameVar string = getPkgName()

// 初始化方法 不能由用户自己调用，也不能被引用, 初始化的的时候自己执行
func init() {
	fmt.Println("pkg1 init method invoked")
}

func getPkgName() string {
	fmt.Println("pkg1.PkgNameVar has been initialized")
	return PkgName
}
