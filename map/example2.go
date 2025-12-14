package _map

import (
	"fmt"
)

type User struct {
	Name string
}

func Test() {
	um := make(map[string]*User, 2)
	fmt.Println(um)
	userList := []User{
		{Name: "Tom"},
		{Name: "Jack"},
	}

	for _, u := range userList {
		um[u.Name] = &u
	}
	fmt.Println(um)
}

func NewArray(size int) {

}
