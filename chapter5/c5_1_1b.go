/*
代码所在章节：5.1.1节
*/

package main

import (
	"runtime"
	"time"
)

func sum() {
	sum := 0
	for i := 0; i < 10000; i++ {
		sum += i
	}
	println(sum)
	time.Sleep(1 * time.Second)
}

func main() {
	go sum()
	//NumGoroutine可以返回当前程序的goroutine数目
	println("NumGoroutine=", runtime.NumGoroutine())

	//main goroutine故意sleep5秒,防止其提前退出
	time.Sleep(5 * time.Second)
}
