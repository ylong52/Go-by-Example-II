//Go 匿名函数(go语言的匿名函数就是闭包)

package main

import (
	"fmt"
)

func outer(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func main() {
	f := outer(10)
	fmt.Println(f(100))
}

/*
output:
$ go run anonymous.go
110
*/
