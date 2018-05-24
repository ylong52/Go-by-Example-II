//Go 信道遍历

package main

import (
	"fmt"
)

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}

/*
output:
$ go run range-over-channels.go
one
two
*/
