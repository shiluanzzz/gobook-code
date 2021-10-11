/*
代码所在章节：4.4.2节
*/

//iface.go
package main

type Caler interface {
	Add(a, b int) int
	Sub(a, b int) int
}

type Adder struct{ id int }

//go:noinline
func (adder Adder) Add(a, b int) int { return a + b }

//go:noinline
func (adder Adder) Sub(a, b int) int { return a - b }

func main() {
	var m Caler = Adder{id: 1234}
	m.Add(10, 32)
}
