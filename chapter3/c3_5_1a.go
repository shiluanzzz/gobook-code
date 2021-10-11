/*
代码所在章节：3.5节
*/

package main

import "fmt"

//有名函数定义，函数名是add
//add类型是函数字面量类型 func (int,int) int
func add(a, b int) int {
	return a + b
}

//函数声明语句，用在go代码调用汇编代码
//func add(int,int) int

//add函数的签名，实际上就是add的字面量类型
//func (int,int) int

//新定义函数类型ADD
//ADD底层类型是函数字面量类型func (int,int) int
type ADD func(int, int) int

//add和ADD的底层类型相同，并且add是字面量类型
//所以add可直接赋值给ADD类型的变量f
var g ADD = add

func main() {
	//匿名函数
	//其不能独立存在，常常作为函数参数、返回值或者赋值给某个变量
	//匿名函数可以看作是函数字面量类型直接显式初始化
	//匿名函数的类型也是函数字面量类型func (int,int) int
	f := func(a, b int) int {
		return a + b
	}
	g(1, 2)
	f(1, 2)

	// f 和add的函数签名相同
	fmt.Printf("%T\n", f)   // func(int, int) int
	fmt.Printf("%T\n", add) // func(int, int) int
}
