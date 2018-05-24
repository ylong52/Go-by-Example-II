//Go 超时

package main

import (
	"fmt"
	"time"
)

/*
	timeout对于连接到外部资源或需要限制执行时间的程序很重要。
	在Go中利用select和channel可以非常容易的实现timeout
*/

func main() {

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}

/*
output:
$ go run timeout.go
timeout 1
result2
*/
