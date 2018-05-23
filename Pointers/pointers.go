//Go 指针

package main

/*
	指针声明格式：
	var var_name *var-type
*/

import (
	"fmt"
)

//传入一个整型类型的参数
func zeroval(ival int) {
	ival = 0
}

//传入一个指向整型的指针参数
func zeroptr(iptr *int) {
	// 声明指针的实际变量
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// & 是Go的取地址符，放到一个变量前使用就会返回相应变量的内存地址
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}

/*
output:
$ go run pointers.go
initial: 1
zeroval: 1
zeroptr: 0
pointer: 0xc042060080
*/

/*
	Go 空指针
	当一个指针被定义后没有分配到任何变量时它的值为nil,nil指针也称为空指针
	一个指针变量通常缩写为 ptr
*/
