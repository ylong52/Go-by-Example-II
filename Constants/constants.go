//Go定义常量

package main

import (
	"fmt"
	"math"
)

//const 定义一个常量(全局)
const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	//常量的定义同样可以使用 简短声明方式(:=)
	nn := 500000000
	fmt.Println(nn)

	//科学计算
	const d = 3e20 / n
	fmt.Println(d)

	//类型转换
	fmt.Println(int64(d))

	//函数计算
	fmt.Println(math.Sin(n))
}

/*
output:
$ go run constants.go
constant
500000000
6e+11
600000000000
-0.28470407323754404
*/
