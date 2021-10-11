/*
代码所在章节：2.5.1节
*/

package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	defer func() {
		panic("first defer panic")
	}()

	defer func() {
		panic("second defer panic")
	}()

	panic("main body panic")
}
