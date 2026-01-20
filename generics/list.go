package main

type List[T any] interface {
	Add(idx int, t T)
	Append(t T)
}

func UseList() {
	var l List[int]
	l.Append(12)

	var lany List[any]
	lany.Append(12.3)
	lany.Append(123)
}

type LinkedList[T any] struct {
	head *node[T]
	t    T
}

type node[T any] struct {
	val T
}
