//Go 函数多返回值
package main

import (
	"fmt"
)

func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a, b)
}

/*
output:
$ go run multiple-return-values.go
3 7
*/
