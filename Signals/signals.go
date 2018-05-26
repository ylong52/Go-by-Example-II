//Go 信号

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

/*
	有时候我们希望Go能智能的处理Unix信号。
	例如，我们希望当服务器接收到一个sigterm信号时能改自定关机，或者一个命令行工具在接收到一个sigint信号时停止处理输入信息
*/

func main() {

	//Go通过向一个通道发送os.Signal值来进行信号通知。我们将创建一个通道来接收这些通知
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	//signal.Notify注册这个给定的通道用于接收特定信号
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	//阻塞。直到上面的协程发送的done值才回继续往下执行然后退出
	<-done
	fmt.Println("exiting")
}

/*
output:
$ go run signals.go
awaiting signal

interrupt
exiting
*/
