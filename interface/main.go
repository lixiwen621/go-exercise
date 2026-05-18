package main

func main() {
	in := Inner{}
	in.DoSomething()

	out := Outer{}
	out.DoSomething()
	// 会输出hello, inner
	out.SayHello()

	ReceiverWithInterfaceDemo()
	ReadWriterDemo()
}
