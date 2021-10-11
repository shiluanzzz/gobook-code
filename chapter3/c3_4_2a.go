/*
代码所在章节：3.4.2节
*/

package main

//import "fmt"

type X struct {
	a int
}

type Y struct {
	X
}

type Z struct {
	*X
}

func (x X) Get() int {
	return x.a
}

func (x *X) Set(i int) {
	x.a = i
}

func main() {

	x := X{a: 1}

	y := Y{
		X: x,
	}

	println(y.Get()) // 1

	//此处编译器做了自动转换
	y.Set(2)
	println(y.Get()) // 2

	//为了不让编译器做自动转换, 我们使用method expression格式调用

	(*Y).Set(&y, 3)

	// type Y的方法集合并没有Set这个方法, 所以下一句编译通不过
	//Y.Set(y, 3) // type Y has no method Set
	println(y.Get()) // 3

	z := Z{
		X: &x,
	}

	//按照嵌套字段的方法集的规则
	//Z 内嵌字段*X,所以type Z和type *Z方法集都是 Get和Set
	//为了不让编译器做自动转换, 我们仍然使用method expression格式调用

	Z.Set(z, 4)
	println(z.Get()) // 4

	(*Z).Set(&z, 5)
	println(z.Get()) // 5

}
