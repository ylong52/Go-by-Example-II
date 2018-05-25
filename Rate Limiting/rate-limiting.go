//Go 并发之限速

package main

import (
	"fmt"
	"time"
)

func main() {

	/*===========================================*/
	//创建缓冲大小为5的requests信道
	requests := make(chan int, 5)
	//requests收到5个数据
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	//速度限制，time.Tick(隔一段时间后)，延迟200ms
	limiter := time.Tick(200 * time.Millisecond)

	//遍历requests信道
	for req := range requests {
		//<-limiter这里是阻塞的，利用了time的channel，缓冲200ms
		<-limiter
		//每隔200ms输出一次
		fmt.Println("requests", req, time.Now())
	}

	/*===========================================*/

	//创建一个缓冲大小为3， 存放时间点的并发限制的信道
	//make(chan time.Time)存放时间点的
	burstyLimiter := make(chan time.Time, 3)

	//将时间数据传给burstyLimiter信道
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	//创建一个缓冲大小为5,int类型的并发请求信道
	burstyRequests := make(chan int, 5)
	//并发请求信道收到5个数据
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		//前三次是并发的
		<-burstyLimiter
		fmt.Println("burstyrequests", req, time.Now())
	}
}

/*
output:
$ go run rate-limiting.go
requests 1 2018-05-25 09:42:02.183599 +0800 CST m=+0.203020701
requests 2 2018-05-25 09:42:02.3839144 +0800 CST m=+0.403336101
requests 3 2018-05-25 09:42:02.5838099 +0800 CST m=+0.603231601
requests 4 2018-05-25 09:42:02.7848099 +0800 CST m=+0.804231601
requests 5 2018-05-25 09:42:02.9848064 +0800 CST m=+1.004228101
burstyrequests 1 2018-05-25 09:42:02.9848064 +0800 CST m=+1.004228101
burstyrequests 2 2018-05-25 09:42:02.9848064 +0800 CST m=+1.004228101
burstyrequests 3 2018-05-25 09:42:02.9848064 +0800 CST m=+1.004228101
burstyrequests 4 2018-05-25 09:42:03.185778 +0800 CST m=+1.205199701
burstyrequests 5 2018-05-25 09:42:03.3852696 +0800 CST m=+1.404691301
*/
