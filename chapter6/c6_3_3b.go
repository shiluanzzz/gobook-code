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

func Format(name string, company S1, level S2, age int) {
	fmt.Printf("name=%s, company=%s, level=%s, age=%d!\n", name, company, level, age)
}

func main() {

	//控制实例创建
	inj := inject.New()

	//实参注入
	inj.Map("tom")
	inj.MapTo("tencent", (*S1)(nil))
	inj.MapTo("T4", (*S2)(nil))
	inj.Map(23)

	//反转调用
	inj.Invoke(Format)
}
