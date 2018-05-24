//Go 非阻塞信道操作

package main

import (
	"fmt"
)

/*
	信道上的基本收发处于阻塞状态。但是我们可以使用select一个default子句来实现非阻塞的收发，甚至是非阻塞的多路select。
*/

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	//使用多个select来实现一个多路无阻塞

	select {
	case msg := <-messages:
		fmt.Println("received msg", msg)
	default:
		fmt.Println("No msg")
	}

	msg := "Hi"
	select {
	case messages <- msg:
		fmt.Println("sent msg", msg)
	default:
		fmt.Println("No msg sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received msg", msg)
	case sig := <-signals:
		fmt.Println("recevied signal", sig)
	default:
		fmt.Println("No activity")
	}
	//msg不能发送到messages信道，因为该信道没有缓冲和选择所以直接执行default语句
}

/*
output:
$ go run non-blocking-channel-operations.go
No msg
No msg sent
No activity
*/
