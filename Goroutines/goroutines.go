package main

import (
	"fmt"
	"time"
)

//Go go中的协程goroutine

/*
	Go协程(goroutine)是与其他函数或方法同时运行的函数或方法。
	可以认为Go协程是轻量级的线程。与创建线程相比，创建Go协程的成本很小。因此在Go中同时运行上千个协程是很常见的。

	Go协程的优点：开销小，多路复用，通过信道(channel)通信

	创建一个Go协程: 在函数或方法调用之前加上关键字go就开启了一个并发的Go协程

	当创建一个Go协程时，创建这个Go协程的语句立即返回。与函数不同，程序流程不会等待Go协程结束再继续执行。程序流程在开启Go协程后立即返回并开始执行下一行代码，忽略Go协程的任何返回值。
	在主协程存在时才能运行其他协程，主协程终止则程序终止，其他协程也将终止。

*/

/*
func hello() {
	fmt.Println("Hello")
}

func main() {
	go hello()
	fmt.Println("main func")
}

output:
$ go run goroutines.go
main func

这里只能看到主函数的输出结果，我们开启的协程输出看不到，主函数调用go heelo()之后，
程序的流程直接调转到吓一跳语句执行，并没有等待hello协程退出然后打印"main func"。
接着主协程结束运行不会再执行任何代码，所以hello协程没有得到运行的机会
*/

//修改下上面的代码
/*
func hello() {
	fmt.Println("Hello")
}

func main() {
	go hello()
	time.Sleep(1 * time.Second)
	fmt.Println("main func")
}
output：
$ go run goroutines.go
Hello
main func

主协程休眠的一秒的过程中，hello协程有了足够的时间在主协程退出之前执行
NOTE:但是要注意的是：我们这里使用sleep函数只是为了更直观了解go的协程，这种阻塞主协程的方法是不推荐的，
	 在Go中提供了信道(channel)来阻塞主协程
*/

//开启多个协程

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func main() {
	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated")
}

/*
	上面的程序开启了两个协程。现在这两个协程同时执行。
	numbers 协程最初睡眠 200 毫秒，然后打印 1，接着再次睡眠然后打印2，以此类推，直到打印到 5。
	类似地，alphabets 协程打印从 a 到 e 的字母，每个字母之间相隔 300 毫秒。
	主协程开启 numbers 和 alphabets 协程，等待 3000 毫秒，最后终止。

	output:
	$ go run goroutines.go
	1 a 2 b 3 4 c 5 d e main terminated
*/
