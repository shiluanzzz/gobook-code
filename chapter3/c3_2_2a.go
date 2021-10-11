/*
代码所在章节：3.2.2节
*/

package main

import (
	"fmt"
)

type MyInt int

type person struct {
	name string
	age  int
	int
}

func (p *person) print() {
	fmt.Print("hello")
}

type SliceInt []int

func (s SliceInt) Sum() int {
	sum := 0
	for _, i := range s {
		sum += i
	}
	return sum
}
func SliceInt_Sum(s SliceInt) int {
	sum := 0
	for _, i := range s {
		sum += i
	}
	return sum
}

func main() {
	var a MyInt = 10
	var b MyInt = 10
	c := a + b
	d := a * b

	fmt.Printf("%d\n", c)
	fmt.Printf("%d\n", d)

	p := person{"tata", 1, 2}
	fmt.Printf("%v\n", p)

	var s SliceInt = []int{1, 2, 3, 4}

	fmt.Println(s.Sum())
	fmt.Println(SliceInt_Sum(s))

}
