/*
代码所在章节：5.1.3节
*/

package main

import (
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var urls = []string{
	"http://www.golang.org/",
	"http://www.google.com/",
	"http://www.qq.com/",
}

func main() {
	for _, url := range urls {

		//每一个url启动一个goroutine，同时给wg加1
		wg.Add(1)

		// Launch a goroutine to fetch the URL.
		go func(url string) {
			//当前goroutine结束后给wg计数减1 ,wg.Done()等价于wg.Add(-1)
			//defer wg.Add(-1)
			defer wg.Done()

			//发卡http get请求并打印http返回码
			resp, err := http.Get(url)
			if err == nil {
				println(resp.Status)
			}
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
}
