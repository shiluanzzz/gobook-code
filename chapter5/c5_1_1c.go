/*
代码所在章节：5.1.1节
*/

package main

import "runtime"

func main() {
	//获取当前GOMAXPROCS的值
	println("GOMAXPROCS=", runtime.GOMAXPROCS(0))

	//设置GOMAXPROCS的值为2
	runtime.GOMAXPROCS(2)

	//获取当前GOMAXPROCS的值
	println("GOMAXPROCS=", runtime.GOMAXPROCS(0))
}
