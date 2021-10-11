/*
代码所在章节：3.1.1节
*/

package main

import "fmt"

//使用type声明的是命名类型
type Person struct {
	name string
	age  int
}

func main() {
	//struct字面量声明的是匿名类型
	a := struct {
		name string
		age  int
	}{"andes", 18}

	fmt.Printf("%T\n", a)
	fmt.Printf("%v\n", a)

	b := Person{"tom", 21}
	fmt.Printf("%T\n", b)
	fmt.Printf("%v\n", b)

}
