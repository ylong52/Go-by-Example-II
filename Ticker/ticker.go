//Go 循环定时器

package main

import (
	"fmt"
	"time"
)

/*
	循环定时器在每次到达指定时间时都触发事件。
*/

func main() {

	//ticker与timer机制类似，都是一个发送值得信道
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		//这里用channel内置的range遍历每500ms接收到的值
		for t := range ticker.C {
			fmt.Println("Ticker at ", t)
		}
	}()

	//ticker可以手动停止(Stop)，一旦停止将不会在其信道上接收到任何信息。
	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

/*
output：
$ go run ticker.go
Ticker at  2018-05-24 22:46:39.062703 +0800 CST m=+0.508029001
Ticker at  2018-05-24 22:46:39.5627316 +0800 CST m=+1.008057601
Ticker at  2018-05-24 22:46:40.0627602 +0800 CST m=+1.508086201
Ticker stopped
*/
