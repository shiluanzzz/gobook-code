/*
代码所在章节：6.1.1节
*/

package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string "学生姓名"
	Age  int    `a:"1111"b:"3333"` //这个不是单引号，而是~键上的符号
}

func main() {
	s := Student{}
	rt := reflect.TypeOf(s)
	fieldName, ok := rt.FieldByName("Name")

	//取tag数据
	if ok {
		fmt.Println(fieldName.Tag)
	}
	fieldAge, ok2 := rt.FieldByName("Age")

	/*可以你JSON 一样，取TAG 里的数据，注意，设置时，两个之间无逗
	号,键名无引号*/
	if ok2 {
		fmt.Println(fieldAge.Tag.Get("a"))
		fmt.Println(fieldAge.Tag.Get("b"))
	}

	fmt.Println("type_Name:", rt.Name())
	fmt.Println("type_NumField:", rt.NumField())
	fmt.Println("type_PkgPath:", rt.PkgPath())
	fmt.Println("type_String:", rt.String())

	fmt.Println("type.Kind.String:", rt.Kind().String())
	fmt.Println("type.String()=", rt.String())

	//获取结构类型那个的字段名称
	for i := 0; i < rt.NumField(); i++ {
		fmt.Printf("type.Field[%d].Name:=%v \n", i, rt.Field(i).Name)
	}

	sc := make([]int, 10)
	sc = append(sc, 1, 2, 3)
	sct := reflect.TypeOf(sc)

	//获取slice元素的Type
	scet := sct.Elem()

	fmt.Println("slice element type.Kind()=", scet.Kind())
	fmt.Printf("slice element type.Kind()=%d\n", scet.Kind())
	fmt.Println("slice element type.String()=", scet.String())

	fmt.Println("slice element type.Name()=", scet.Name())
	fmt.Println("slice type.NumMethod()=", scet.NumMethod())
	fmt.Println("slice type.PkgPath()=", scet.PkgPath())
	fmt.Println("slice type.PkgPath()=", sct.PkgPath())

}
