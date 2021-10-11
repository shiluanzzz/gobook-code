/*
代码所在章节：2.2.1节
*/

package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

type Op func(int, int) int

func do(f Op, a, b int) int {
	return f(a, b)
}

func main() {

	a := do(add, 1, 2)
	fmt.Println(a)

	s := do(sub, 1, 2)
	fmt.Println(s)
}
