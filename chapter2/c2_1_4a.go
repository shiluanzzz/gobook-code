/*
代码所在章节：2.1.4节
*/

package main

import "fmt"

func add(a, b int) (sum int) {
	sum = a + b
	return
}

/*
func add2(a, b int) int {
	return a + b
}


func add(a, b int) (sum int) {
	anonymous := func(x, y int) int {
		return x + y
	}
	return anonymous(a, b)
}
*/

func sum(arr ...int) (sum int) {
	for v := range arr {
		sum += v
	}
	return
}

func suma(arr ...int) (sum int) {
	for v := range arr {
		sum += v
	}
	return
}

func sumb(arr []int) (sum int) {
	for v := range arr {
		sum += v
	}
	return
}

type sumFunc func(...int) int

func do(f sumFunc, a, b int) int {
	return f(a, b)
}

func main() {
	slice := []int{1, 2, 3, 4}
	sum(slice...)

	do(suma, 1, 2)
	fmt.Printf("%T\n", suma) // func(...int) int
	fmt.Printf("%T\n", sumb) // func([]int) int

}
