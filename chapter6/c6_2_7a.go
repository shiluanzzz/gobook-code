/*
代码所在章节：6.2.7节
*/

package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func main() {
	u := User{Id: 1, Name: "andes", Age: 20}

	va := reflect.ValueOf(u)
	vb := reflect.ValueOf(&u)

	//值类型是可修改的
	fmt.Println(va.CanSet(), va.FieldByName("Name").CanSet()) //false false

	//指针类型是可修改的
	fmt.Println(vb.CanSet(), vb.Elem().FieldByName("Name").CanSet()) //false false

	fmt.Printf("%v\n", vb)
	name := "shine"
	vc := reflect.ValueOf(name)

	//通过Set函数修改变量的值
	vb.Elem().FieldByName("Name").Set(vc)
	fmt.Printf("%v\n", vb)

}
