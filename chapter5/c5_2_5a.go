/*
代码所在章节：5.2.5节
*/

package main

import (
	"fmt"
	"time"
)

//一个查询结构体
//这里的sql和result是一个简单的抽象，具体的应用，可能是更复杂的数据类型
type query struct {
	//参数Channel
	sql chan string

	//结果Channel
	result chan string
}

//执行Query

func execQuery(q query) {

	//启动协程
	go func() {
		//获取输入
		sql := <-q.sql

		//访问数据库

		//输出结果通道
		q.result <- "result from " + sql
	}()

}

func main() {

	//初始化Query
	q := query{make(chan string, 1), make(chan string, 1)}

	//执行Query，注意执行的时候无需准备参数
	go execQuery(q)
	//准备参数
	q.sql <- "select * from table;"

	//do otherthings
	time.Sleep(1 * time.Second)

	//获取结果
	fmt.Println(<-q.result)
}
