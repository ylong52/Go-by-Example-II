//Go 关闭信道

package main

import (
	"fmt"
)

/*
	关闭信道表示不会再发送任何值，这对于将完成情况转达给信道的接收器很有用
*/

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			//协程开启的时候jobs信道是阻塞的，它在等待接收数据
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	//给jobs信道传数据，传完后关闭
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	//同步机制来阻塞，等待jobs执行完
	<-done
}

/*
output:
$ go run closing-channels.go
sent job 1
received job 1
sent job 2
received job 2
sent job 3
received job 3
sent all jobs
received all jobs
*/
