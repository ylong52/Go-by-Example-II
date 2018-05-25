//Go 互斥锁

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

/*
	原子操作管理简单的状态，对于复杂的状态我们可以使用互斥锁来跨多个协程安全地访问数据
*/

func main() {

	//建立map型的state
	var state = make(map[int]int)
	//mutex 同步访问state
	var mutex = &sync.Mutex{}
	//记录读写次数
	var readOps uint64
	var writeOps uint64

	//开启100个协程对state重复读取，每个协程每毫秒执行一次
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				//使用锁确保mutex单独访问sate
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				//每执行一次，读操作次数增加一次
				atomic.AddUint64(&readOps, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}

	//开启10个协程模拟写入。与读取的操作模式相同
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}
	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}

/*
output:
$ go run mutexes.go
readOps: 18830
writeOps: 1890
state: map[0:42 2:46 3:78 1:10 4:23]
*/
