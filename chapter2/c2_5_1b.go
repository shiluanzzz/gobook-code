/*
代码所在章节：2.5.1节
*/

package main

import (
	"fmt"
	"time"
)

func do() {
	//do函数里无法捕捉 go da()抛出的异常
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	go da()
	go db()
	time.Sleep(3 * time.Second)
}

func da() {
	panic("panic da")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func db() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func main() {
	do()
}
