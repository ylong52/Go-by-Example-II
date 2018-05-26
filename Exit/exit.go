//Go 退出

package main

import (
	"fmt"
	"os"
)

/*
	使用os.Exit来立即进行带给定状态的退出
*/

func main() {

	//当使用os.Exit时defer将不会被执行
	defer fmt.Println("!")

	//退出状态为3
	os.Exit(3)
}

/*
output:
$ go run exit.go
exit status 3
*/
