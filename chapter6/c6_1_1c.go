/*
代码所在章节：6.1.1节
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

func (this User) String() {
	println("User:", this.Id, this.Name, this.Age)
}

func Info(o interface{}) {
	//获取Value信息
	v := reflect.ValueOf(o)

	//通过Vlue获取Type
	t := v.Type()

	//类型名称
	println("Type:", t.Name())

	//访问接口字段名，字段类型和字段值信息
	println("Fields:")
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()

		//类型查询
		switch value := value.(type) {
		case int:
			fmt.Printf(" %6s: %v = %d\n", field.Name, field.Type, value)
		case string:
			fmt.Printf(" %6s: %v = %s\n", field.Name, field.Type, value)
		default:
			fmt.Printf(" %6s: %v = %s\n", field.Name, field.Type, value)

		}
	}
}
func main() {
	u := User{1, "Tom", 30}
	Info(u)
}
