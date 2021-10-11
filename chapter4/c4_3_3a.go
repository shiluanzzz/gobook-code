/*
代码所在章节：4.3.3节
*/

package main

import "fmt"

type Inter interface {
	Ping()
	Pang()
}

type St struct{}

func (St) Ping() {
	println("ping")
}
func (*St) Pang() {
	println("pang")
}

func main() {
	var st *St = nil
	var it Inter = st

	fmt.Printf("%p\n", st)
	fmt.Printf("%p\n", it)

	if it != nil {
		it.Pang()

		//下面的语句会导致panic
		//panic: value method main.St.Ping called using nil *St pointer
		//方法转换为函数调用，第一个参数是St类型，由于*St是nil，无法获取指针所指的对象值，所以panic.
		//it.Ping()
	}
}
