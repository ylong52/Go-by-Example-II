//Go 随机数

package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	Go的math/rand包提供了伪随机数生成器
*/

func main() {
	//rand.Intn返回一个随机的整数
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	//返回一个随机的64位浮点数
	fmt.Println(rand.Float64())

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()
}

/*
output:
$ go run random-numbers.go
81,87
0.6645600532184904
78,43
*/
