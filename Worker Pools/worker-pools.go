//Go 并发之工作者池

package main

import (
	"fmt"
	"time"
)

//建立一个工作者函数。参数：id, 分发任务的jobs信道，收集工作结果的results信道
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		//完成后将结果传给results信道
		results <- j * 2
	}
}

func main() {
	//建立两个缓冲大小为100的信道
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	//创建3个工作者
	for w := 1; w <= 3; w++ {
		//开启工作者协程，工作者同时工作实现简单的并发
		go worker(w, jobs, results)
	}

	//创建5个任务并传给jobs信道
	for j := 1; j <= 5; j++ {
		fmt.Println("make job：", j)
		jobs <- j
	}
	close(jobs)

	//这里创建了和任务数量相等的收集结果的信道
	/*
		我们可以试着将a改成 a<5,此时会出现dead lock。
		为什么呢？工作者完成每个任务后会把结果发送给results信道，当完成第4个任务时，
		此时没有results信道可用，做完第五个任务时需要将结果发送给results信道，但是没有信道
		可以接收这个结果，导致他一直等待信道来接收它，这样就触发了Go的自我保护机制，自己kill了。
		丢出error=>  fatal error: all goroutines are asleep - deadlock!
	*/
	for a := 1; a <= 5; a++ {
		<-results
	}
}

/*
output:
$ go run worker-pools.go
make job： 1
make job： 2
make job： 3
make job： 4
make job： 5
worker 3 started job 3
worker 2 started job 2
worker 1 started job 1
worker 2 finished job 2
worker 2 started job 4
worker 1 finished job 1
worker 1 started job 5
worker 3 finished job 3
worker 1 finished job 5
worker 2 finished job 4
*/
