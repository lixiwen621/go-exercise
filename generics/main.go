package main

import "fmt"

func main() {
	maxNum, _ := getMax(10, 5, 6, 8)
	fmt.Println(maxNum)

	minNum, _ := getMin(3, 6, 2, 1)
	fmt.Println(minNum)

	s1 := []int{1, 2, 3, 4, 5}
	newSlice, _ := AddSlice(s1, 5, 7)
	fmt.Println(newSlice)
}
