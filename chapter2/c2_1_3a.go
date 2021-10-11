/*
代码所在章节：2.1.3节
*/

package main

import (
	"fmt"
)

func chvalue(a int) int {
	a = a + 1
	return a
}

func chpointer(a *int) {
	*a = *a + 1
	return
}

func main() {
	a := 10
	chvalue(a)
	fmt.Println(a)
	chpointer(&a)
	fmt.Println(a)

}
