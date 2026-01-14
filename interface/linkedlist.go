package main

type LinkedList struct {
	head *node
	tail *node

	// 这个就是包外可以访问
	Len int
}

func (l LinkedList) Add(idx int, val any) {

}

// 方法接口器, receiver
func (l *LinkedList) AddV1(idx int, val any) {

}

type node struct {
	prev *node
	next *node
}
