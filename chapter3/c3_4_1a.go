/*
代码所在章节：3.4.1节
*/

package main

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
	//z.a, z.Y.a, z.Y.X.a 三者是等价的, z.a z.Y.a是z.Y.X.a的简写
	println(z.a, z.Y.a, z.Y.X.a) //1 1 1

	z = Z{}
	z.a = 2
	println(z.a, z.Y.a, z.Y.X.a) //2 2 2
}
