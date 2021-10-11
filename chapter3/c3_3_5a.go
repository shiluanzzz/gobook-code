/*
代码所在章节：3.3.5节
*/

package main

type Data struct{}

func (Data) TestValue()    {}
func (*Data) TestPointer() {}

func main() {

	var p *Data = nil //p的类型是*Data,但是p的值是nil,p指向的是一块未初始化的类型实例

	//指针方法集的调用是可以使用nil传递进去的,只要方法内部不去访问nil指向的实例，不会引发panic

	//直接调用
	p.TestPointer()            // method value
	(*Data)(nil).TestPointer() // method value 等价p.TestPointer()

	//method value
	m := p.TestPointer // method value
	m()
	m = (*Data)(nil).TestPointer // method value 等价p.TestPointer()
	m()

	//(*p).TestValue() // p为nil，无法通过*p访问p指向的实例，所以此调用会导致panic
	//(Data)(nil).TestValue() // cannot convert nil to type Data
	// p.TestValue() // invalid memory address or nil pointer dereference
	// Data.TestValue(nil) // cannot use nil as type Data in function argument

	//method expression

	(*Data).TestPointer(nil)
	(*Data).TestPointer(p)

	//等价上面调用
	s := (*Data).TestPointer
	s(nil)
	s(p)

	//这种字面量显式调用,编译器不会进行方法集的自动转换,编译器会按照严格的方法集合校验
	// *Data 方法集是TestPointer和TestValue
	(*Data)(&struct{}{}).TestPointer() //显式的调用
	(*Data)(&struct{}{}).TestValue()   //显式的调用

	// Data 方法集只有TestValue
	(Data)(struct{}{}).TestValue() //method value
	Data.TestValue(struct{}{})     //method expression

	//Data.TestPoiter(struct{}{})         //type Data has no method TestPoiter
	//(Data)(struct{}{}).TestPointer()   //cannot call pointer method on Data(struct {} literal)

	var a Data = struct{}{}

	//method expression 调用编译器不会进行自动转换
	Data.TestValue(a)
	//Data.TestValue(&a)
	(*Data).TestPointer(&a)
	//Data.TestPointer(&a)

	//method value 调用编译器不会进行自动转换
	f := a.TestValue
	f()

	y := (&a).TestValue //编译器帮助转换a.TestValue
	y()

	g := a.TestPointer //会转换为(&a).TestPointer
	g()

	x := (&a).TestPointer
	x()

}
