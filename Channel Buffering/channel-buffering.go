//Go 信道缓冲

package main

import (
	"fmt"
)

func main() {

	//创建缓冲信道，缓冲大小为2
	//make(chan string)或make(chan string, 0) 这种是没有缓冲的信道
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

/*
output:
$ go run channel-buffering.go
buffered
channel
*/
