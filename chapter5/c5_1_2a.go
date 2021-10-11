/*
代码所在章节：5.1.2节
*/

package main

import (
	"runtime"
	//	"time"
)

func main() {
	c := make(chan struct{})
	go func(i chan struct{}) {
		sum := 0
		for i := 0; i < 10000; i++ {
			sum += i
		}
		println(sum)
		//写通道
		c <- struct{}{}
	}(c)

	//NumGoroutine可以返回当前程序的goroutine数目
	println("NumGoroutine=", runtime.NumGoroutine())

	//读通道c，通过通道进行同步等待
	<-c
}
