package main

import "fmt"

func main() {
	Func43()
	result := DeferReturn()
	fmt.Println(result)

	result2 := DeferReturnV1()
	fmt.Println(result2)
}
