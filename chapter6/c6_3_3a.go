/*
代码所在章节：6.3.3节
*/

package main

func f1() {
	println("f1")
}

func f2() {
	println("f2")
}

func main() {
	funcs := make(map[string]func())
	funcs["f1"] = f1
	funcs["f2"] = f1

	funcs["f1"]()
	funcs["f2"]()
}
