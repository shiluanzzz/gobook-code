/*
代码所在章节：4.4.1节
*/

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var str interface{} = "Hello World!"
	p := (*struct {
		tab  uintptr
		data uintptr
	})(unsafe.Pointer(&str))

	fmt.Printf("%+v\n", p)
	fmt.Printf("%v\n", p.data)
}
