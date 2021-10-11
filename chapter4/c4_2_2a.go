/*
代码所在章节：4.2.2节
*/

package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var i io.Reader
	switch v := i.(type) { //此处i是为未初始化的接口变量，所以v为nil
	case nil:
		fmt.Printf("%T\n", v) //<nil>
	default:
		fmt.Printf("default")
	}

	f, err := os.OpenFile("notes.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	i = f

	switch v := i.(type) {
	//i的绑定的实例是*osFile类型，实现了io.ReadWriter接口，所以下面case匹配成功。
	case io.ReadWriter:
		//v是io.ReadWriter接口类型，所以可以调用Write方法
		v.Write([]byte("io.ReadWriter\n"))
		fmt.Println("Type Switch Result: io.ReadWriter")

	//由于上一个case已经匹配，就算这个case也匹配，也不会走到这里
	case *os.File:
		v.Write([]byte("*os.File\n"))
		fmt.Println("Type Switch Result: *os.File")
		//这里可以调用具体类型方法
		v.Sync()

	default:
		fmt.Println("Type Switch Result: unknown")
		return
	}

	switch v := i.(type) {
	//匹配成功，v的类型就是具体类型*os.File
	case *os.File:
		v.Write([]byte("*os.File\n"))
		fmt.Println("Type Switch Result: *os.File")
		v.Sync()

	//由于上一个case已经匹配，就算这个case也匹配，也不会走到这里
	case io.ReadWriter:
		//v是io.ReadWriter接口类型，所以可以调用Write方法
		v.Write([]byte("io.ReadWriter\n"))
		fmt.Println("Type Switch Result: io.ReadWriter")

	default:
		fmt.Println("Type Switch Result: unknown")
		return
	}

	switch v := i.(type) {
	//多个类型，f满足其中任何一个就算匹配
	case *os.File, io.ReadWriter:
		// 此时相当于执行力v := i ,v和i是等价的，使用v没有意义。
		if v == i {
			fmt.Println(true) //true
		}
	default:
		return
	}

}
