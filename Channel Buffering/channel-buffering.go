//Go 信道缓冲

package main

import (
	"fmt"
)

func main() {
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
