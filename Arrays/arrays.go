//Go 数组

package main

import (
	"fmt"
)

func main() {

	//定义一个长度为5的数组(未初始化，每个元素为0)
	var a [5]int
	fmt.Println("emp:", a)

	//给a数组的第5个元素赋值
	a[4] = 100
	//输出新的数组
	fmt.Println("set:", a)
	//获取a数组的第5个元素
	fmt.Println("get:", a[4])
	//获取a数组的长度
	fmt.Println("len:", len(a))

	//定义一个长度为5的数组，并初始化
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	//定义一个2维数组，长度为3(未初始化，每个元素为0)
	var twoD [2][3]int
	fmt.Println("twoD:", twoD)
	//两个for循环初始化这个二维数组
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("new_twoD:", twoD)
}

/*
output:
$ go run arrays.go
emp: [0 0 0 0 0]
set: [0 0 0 0 100]
get: 100
len: 5
dcl: [1 2 3 4 5]
twoD: [[0 0 0] [0 0 0]]
new_twoD: [[0 1 2] [1 2 3]]
*/
