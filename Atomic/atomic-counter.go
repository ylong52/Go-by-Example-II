//Go atomic原子操作

package main

/*
	原子操作共有5种:增减、比较并交换、载入、存储和交换
*/

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var ops uint64
	//启动50个协程
	for i := 0; i < 50; i++ {
		go func() {
			for {
				//原子增操作
				//给计数器分配内存“&”
				atomic.AddUint64(&ops, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	time.Sleep(time.Second)

	//原子载入
	//从内存中复制一个出来opsFinal
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}

/*
output:
$ go run atomic-counter.go
ops: 31993 //执行了30000多次
*/
