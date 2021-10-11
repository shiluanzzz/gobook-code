/*
代码所在章节：6.3.3节
*/

package main

import (
	"fmt"
	"github.com/codegangsta/inject"
)

type S1 interface{}
type S2 interface{}

type Staff struct {
	Name    string `inject`
	Company S1     `inject`
	Level   S2     `inject`
	Age     int    `inject`
}

func main() {
	//创建被注入实例
	s := Staff{}

	//控制实例创建
	inj := inject.New()

	//初始化注入值
	inj.Map("tom")
	inj.MapTo("tencent", (*S1)(nil))
	inj.MapTo("T4", (*S2)(nil))
	inj.Map(23)

	//实现注入
	inj.Apply(&s)

	//打印结果
	fmt.Printf("s=%v\n", s)

}
