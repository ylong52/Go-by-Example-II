//Go 标准定时器

package main

import (
	"fmt"
	"time"
)

/*
	Go的标准包time里提供了两种定时器，一种是标准定时器Timer，另一种是循环定时器Ticker
	Timer标准定时器在到达指定时间时仅触发一次。
*/

func main() {

	//新建一个标准计时器，两秒以后触发(Go触发计时器的方法就是在计时器的channel中发送值)
	timer1 := time.NewTimer(2 * time.Second)
	//此处等待channel中的信号，阻塞2秒
	<-timer1.C
	fmt.Println("Timer 1 expired")

	//新建计时器，一秒后触发
	timer2 := time.NewTimer(time.Second)
	//新开启一个协程来处理触发后的事件
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	//上面的协程在等待信号是阻塞的，执行到这里的时候直接把定时器停掉了，所以上面的协程是没有执行完的
	stop2 := timer2.Stop()

	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}

/*
output：
$ go run timer.go
Timer 1 expired
Timer 2 stopped
*/
