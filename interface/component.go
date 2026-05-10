package main

type Inner struct {
}

func (in Inner) DoSomething() {

}

type Outer struct {
	Inner
}

type OuterPtr struct {
	*Inner
}

type Component struct {
	Outer
}

func (o Outer) Name() string {
	return "Outer"
}

func (i Inner) SayHello() {
	println("hello," + i.Name())
}

func (i Inner) Name() string {
	return "Inner"
}
