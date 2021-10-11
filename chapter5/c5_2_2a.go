/*
代码所在章节：5.2.2节
*/

package main

import (
	"fmt"
)

//chain函数输入参数和输入参数类型相同都是chan int类型
//chain函数功能是将chan内的数据统一加1

func chain(in chan int) chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- 1 + v
		}
		close(out)
	}()
	return out
}

func main() {
	in := make(chan int)

	//初始化输入参数
	go func() {
		for i := 0; i < 10; i++ {
			in <- i
		}
		close(in)
	}()

	//连续调用3次chan，想当与把in中的每个元素都加3
	out := chain(chain(chain(in)))
	for v := range out {
		fmt.Println(v)
	}
}
