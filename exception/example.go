package main

import (
	"errors"
	"fmt"
)

func TestError(numA int, numB int) (result int, err error) {
	if numB == 0 {
		err = errors.New("除数不能为0")
		return 0, err
	}
	result = numA / numB
	return
}

func Main() {
	num, err := TestError(10, 0)
	if err != nil {
		fmt.Println(err)
		fmt.Println(num)
	} else {
		fmt.Println(num)
	}
}
