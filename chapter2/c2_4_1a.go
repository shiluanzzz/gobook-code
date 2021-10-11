/*
代码所在章节：2.4.1节
*/

package main

func fa(a int) func(i int) int {

	return func(i int) int {
		println(&a, a)
		a = a + i
		return a
	}
}

func main() {
	f := fa(1) //f引用的外部的闭包环境包括本次函数调用的形参a的值1
	g := fa(1) //f引用的外部的闭包环境包括本次函数调用的形参a的值1
	// 此时f,g引用的闭包环境中的a的值并不是同一个，而是两次函数调用产生的副本

	println(f(1)) //2
	//多次调用f引用的是同一个副本a,所以
	println(f(1)) //3

	//g中a的值仍然是1
	println(g(1)) //2
	println(g(1)) //3
}
