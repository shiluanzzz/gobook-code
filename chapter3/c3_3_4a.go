/*
代码所在章节：3.3.4节
*/

package main

import "fmt"

type Int int

func (a Int) Max(b Int) Int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func (i *Int) Set(a Int) {
	*i = a
}

func (i Int) Print() {
	fmt.Printf("value=%d\n", i)
}

func main() {
	var a Int = 10
	var b Int = 50

	c := a.Max(b)
	c.Print()
	(&a).Print()

	a.Set(20)
	a.Print()

	(&a).Set(30)
	a.Print()

}
