/*
代码所在章节：6.1.1节
*/

package main

import (
	"reflect"
)

type INT int
type A struct {
	a int
}
type B struct {
	b string
}
type Ita interface {
	String() string
}

func (b B) String() string {
	return b.b
}
func main() {
	var a INT = 12
	var b int = 14

	//对于实参是具体类型,reflect.TypeOf返回是其静态类型
	ta := reflect.TypeOf(a)
	tb := reflect.TypeOf(b)

	//INT 和int是两个类型，二者不相等
	if ta == tb {
		println("ta==tb")
	} else {
		println("ta!=tb") //ta!=tb
	}
	println(ta.Name()) //INT
	println(tb.Name()) // int

	//底层基础类型
	println(ta.Kind().String()) //int
	println(tb.Kind().String()) //int
	s1 := A{1}
	s2 := B{"tata"}

	//对于实参是具体类型,reflect.TypeOf返回是其静态类型
	println(reflect.TypeOf(s1).Name()) //A
	println(reflect.TypeOf(s2).Name()) //B

	//Type的Kind()方法返回的是基础类型,类型A和B的底层基础类型都是struct
	println(reflect.TypeOf(s1).Kind().String()) //struct
	println(reflect.TypeOf(s2).Kind().String()) //struct
	ita := new(Ita)
	var itb Ita = s2

	//对于实参是未绑定具体变量的接口类型,reflect.TypeOf返回的是接口类型本身
	//也就是接口的静态类型
	println(reflect.TypeOf(ita).Elem().Name())          //Ita
	println(reflect.TypeOf(ita).Elem().Kind().String()) //interface

	//对于实参是绑定了具体变量的接口类型,reflect.TypeOf返回的是绑定的具体类型
	//也就是接口的动态类型
	println(reflect.TypeOf(itb).Name())          //B
	println(reflect.TypeOf(itb).Kind().String()) //struct
}
