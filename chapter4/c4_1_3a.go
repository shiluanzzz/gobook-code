/*
代码所在章节：4.1.3节
*/

package main

type Printer interface {
	Print()
}

type S struct{}

func (s S) Print() {
	println("print")
}

func main() {
	var i Printer
	//	i.Print() // panic: runtime error: invalid memory address or nil pointer dereference

	//必须初始化
	i = S{}
	i.Print()

}
