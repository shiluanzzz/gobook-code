/*
代码所在章节：2.4.1节
*/

package main

func fa(base int) (func(int) int, func(int) int) {
	println(&base, base)
	add := func(i int) int {
		base += i
		println(&base, base)
		return base
	}
	sub := func(i int) int {
		base -= i
		println(&base, base)
		return base
	}
	return add, sub
}
func main() {
	//f,g 闭包引用的base是同一个，是fa函数调用是的形参
	f, g := fa(0)
	//s,k 闭包引用的base是同一个，是fa函数调用的的形参
	s, k := fa(0)
	//f,g 和 s,k引用不同的闭包变量，这个是由于fa每次调用都要重新分配形参
	println(f(1), g(2))
	println(s(1), k(2))
}
