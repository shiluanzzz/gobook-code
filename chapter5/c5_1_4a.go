/*
代码所在章节：5.1.4节
*/

package main

func main() {
	ch := make(chan int, 1)
	go func(chan int) {
		for {
			select {
			//0或者1写入是随机的
			case ch <- 0:
			case ch <- 1:
			}
		}
	}(ch)
	for i := 0; i < 10; i++ {
		println(<-ch)
	}

}
