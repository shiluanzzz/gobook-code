/*
代码所在章节：3.4.1节
*/

package main

type X struct {
	a int
}
type Y struct {
	X
	a int
}
type Z struct {
	Y
	a int
}

func main() {
	x := X{a: 1}
	y := Y{
		X: x,
		a: 2,
	}
	z := Z{
		Y: y,
		a: 3,
	}
	//此时的z.a, z.Y.a, z.Y.X.a 代表不同的字段
	println(z.a, z.Y.a, z.Y.X.a) // 3 2 1
	z = Z{}
	z.a = 4
	z.Y.a = 5
	z.Y.X.a = 6
	//此时的z.a, z.Y.a, z.Y.X.a 代表不同的字段
	println(z.a, z.Y.a, z.Y.X.a) // 4 5 6
}
