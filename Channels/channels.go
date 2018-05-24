//Go 信道

package main

import (
	"fmt"
)

/*
	信道：Go协程之间通信的管道

	信道声明方式：
	下面定义了一个int型的信道
	一：
		var name chan type
		name = make(chan int)
	二：
		name := make(chan int)

	信道收发数据的语法：
	data := <- a   //read from channel a
	a <- data 	  //write to channel a
	箭头的指向说明了数据是收还是发

	通过信道的收发数据默认是阻塞的:
	当数据发送给信道后，程序流程在发送语句处阻塞，直到其他协程从该信道中读取数据。
	同样地，当从信道读取数据时，程序在读取语句处阻塞，直到其他协程发送数据给该信道。
	信道的这种特性使得协程间通信变得高效，而不是向其他编程语言一样，显式的使用锁和条件变量来达到此目的。
*/

/*
func hello(done chan bool) {
	fmt.Println("Hello")
	done <- true
}
func main() {
	//定义了一个bool类型的信道
	done := make(chan bool)
	//将done信道作为参数传递给hello()协程
	go hello(done)
	//从信道中读取数据，但我们没有使用该数据也没有将它赋值给其他变量。
	//此时主协程被阻塞，等待从信道done中读取数据
	<-done
	//hello()协程接收信道done，执行完后将数据写入done信道，写入完毕后，主协程从信道中接收到数据，
	//此时主协程继续执行
	fmt.Println("main func")
}

output:
$ go run channels.go
Hello
main func
*/

//计算平方和与立方和(开启两个协程)，并求他们的总和(主协程)

//calSquares 计算平方和
func calSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

//calCubes 计算立方和
func calCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

func main() {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	//将平方和计算结果发送给了信道sqrch
	go calSquares(number, sqrch)
	//将立方和计算结果发送给了信道cubech
	go calCubes(number, cubech)
	//主协程收到数据继续执行
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
}

/*
output：
$ go run channels.go
Final output 1536
*/
