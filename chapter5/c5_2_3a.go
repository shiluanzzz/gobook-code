/*
代码所在章节：5.2.3节
*/

package main

import (
	"fmt"
)

//工作池的goroutine数目
const (
	NUMBER = 10
)

//工作任务
type task struct {
	begin  int
	end    int
	result chan<- int
}

//任务处理:计算begin到end的和
//执行结果写入到结果chan result中
func (t *task) do() {
	sum := 0
	for i := t.begin; i <= t.end; i++ {
		sum += i
	}
	t.result <- sum
}

func main() {
	workers := NUMBER

	//工作通道
	taskchan := make(chan task, 10)

	//结果通道
	resultchan := make(chan int, 10)

	//worker信号通道
	done := make(chan struct{}, 10)

	//初始化task的goroutine,计算1000个自然数之和
	go InitTask(taskchan, resultchan, 1000)

	//分发任务在NUMBER个goroutine 池
	DistributeTask(taskchan, workers, done)

	//获取各个goroutine处理完任务的通知，并关闭结果通道
	go CloseResult(done, resultchan, workers)

	//通过结果通道处理结果
	sum := ProcessResult(resultchan)

	fmt.Println("sum=", sum)
}

//初始化待处理task chan
func InitTask(taskchan chan<- task, r chan int, p int) {
	qu := p / 10
	mod := p % 10
	high := qu * 10
	for j := 0; j < qu; j++ {
		b := 10*j + 1
		e := 10 * (j + 1)
		tsk := task{
			begin:  b,
			end:    e,
			result: r,
		}
		taskchan <- tsk
	}
	if mod != 0 {
		tsk := task{
			begin:  high + 1,
			end:    p,
			result: r,
		}
		taskchan <- tsk
	}

	close(taskchan)
}

//读取task chan 分发到worker goroutine 处理，workers的总的数量是workers
func DistributeTask(taskchan <-chan task, workers int, done chan struct{}) {

	for i := 0; i < workers; i++ {
		go ProcessTask(taskchan, done)
	}
}

//工作goroutine处理具体工作，并将处理结构发送到结果chan
func ProcessTask(taskchan <-chan task, done chan struct{}) {
	for t := range taskchan {
		t.do()
	}
	done <- struct{}{}
}

//通过done channel来同步等待所有工作goroutine的结束，然后关闭结果chan
func CloseResult(done chan struct{}, resultchan chan int, workers int) {
	for i := 0; i < workers; i++ {
		<-done
	}
	close(done)
	close(resultchan)
}

// 读取结果通道，汇总结果
func ProcessResult(resultchan chan int) int {
	sum := 0
	for r := range resultchan {
		sum += r
	}
	return sum
}
