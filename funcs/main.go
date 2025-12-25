package main

import "fmt"

func main() {
	a := DeferReturn()
	fmt.Println(a)

	b := DeferReturnV1()
	fmt.Println(b)

	DeferConditional(true)
}
