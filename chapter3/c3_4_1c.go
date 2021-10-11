/*
代码所在章节：3.4.1节
*/

package main

import "fmt"

type X struct {
	a int
}
type Y struct {
	X
	b int
}
type Z struct {
	Y
	c int
}

func (x X) Print() {
	fmt.Printf("In X, a=%d\n", x.a)
}
func (x X) XPrint() {
	fmt.Printf("In X, a=%d\n", x.a)
}
func (y Y) Print() {
	fmt.Printf("In Y, b=%d\n", y.b)
}
func (z Z) Print() {
	fmt.Printf("In Z, c=%d \n", z.c)
	//显式的完全路径调用内嵌字段的方法
	z.Y.Print()
	z.Y.X.Print()
}
func main() {
	x := X{a: 1}
	y := Y{
		X: x,
		b: 2,
	}
	z := Z{
		Y: y,
		c: 3,
	}

	//从外向内查找首先找到的是Z的Print()方法
	z.Print()

	//从外向内查找,最后找到的是X的XPrint()方法
	z.XPrint()
	z.Y.XPrint()
}
