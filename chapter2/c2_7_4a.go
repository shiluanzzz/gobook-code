/*
代码所在章节：2.7.4节
*/

package main

func a(i int) func() {
	return func() {
		i++
		print(i)
	}
}

func main() {
	f := a(1)
	f()
}
