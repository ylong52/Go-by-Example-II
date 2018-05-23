//Go的func(函数)

package main

import (
	"fmt"
)

//定义一个函数，传入参数为两个int类型的元素， 返回值是int类型
func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	//定义一个变量
	res := plus(1, 2)
	fmt.Println("1+2=", res)
	//引用变量(必须先定义了才能使用)
	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3=", res)
}

/*
output:
$ go run functions.go
1+2= 3
1+2+3= 6
*/
