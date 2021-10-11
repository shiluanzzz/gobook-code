/*
代码所在章节：3.1.3节
*/

package main

import (
	"fmt"
)

type Map map[string]string

func (m Map) Print() {
	for _, key := range m {
		fmt.Println(key)
	}
}

type iMap Map

func (m iMap) Print() {
	for _, key := range m {
		fmt.Println(key)
	}
}

type slice []int

func (s slice) Print() {
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	mp := make(map[string]string, 10)
	mp["hi"] = "tata"
	var ma Map = mp
	//var im iMap = ma
	//var im iMap = mp
	var im iMap = (iMap)(ma)

	ma.Print()
	im.Print()

	//Map实现了Print()所以其可以赋值给接口变量
	var i interface {
		Print()
	} = ma

	i.Print()
	s1 := []int{1, 2, 3}
	var s2 slice
	s2 = s1
	s2.Print()

	s := "hello,世界!"
	var a []byte
	a = []byte(s)
	var b string
	b = string(a)
	var c []rune
	c = []rune(s)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

}
