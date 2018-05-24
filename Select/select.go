//Go select

package main

import (
	"fmt"
	"time"
)

/*
	select是Go中的一个控制结构，类似于switch语句，用于处理异步IO操作。
	select会监听case语句中channel的读写操作，当case中channel读写操作为非阻塞状态(即能读写)时，将会触发相应动作
	将goroutine、channel、select结合起来是Go的一个强大的功能

	！！！
	select中的case语句必须是一个channel操作
	select中的default子句总是可运行的

	如果有多个case都可以运行，select会随机公平地选出一个执行，其他不会执行。
	如果没有可运行的case语句，且有default语句，那么就会执行default的动作。
	如果没有可运行的case语句，且没有default语句，select将阻塞，直到某个case通信可以运行
*/

func main() {
	//定义两个string类型的通道
	c1 := make(chan string)
	c2 := make(chan string)

	//此协程休眠一秒后 将"one"这个数据传给信道c1
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	//此协程休眠一秒后 将"two"这个数据传给信道c2
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	//上面这两个协程同时运行，由于他们内部休眠时间不同，所以对应的信道收到数据的时间也不同

	for i := 0; i < 3; i++ {
		select {
		//c1信道的数据发送
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		//c2信道的数据发送
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		default:
			fmt.Println("Nothing")
			time.Sleep(2 * time.Second)
		}
		//协程执行但处于休眠状态，主函数进入到default后休眠2秒(此时是阻塞的)，然后c1和c2信道发出数据，执行case，程序结束
	}
}

/*
output:
$ go run select.go
Nothing
received one
received two
*/
