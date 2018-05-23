//Go定义变量

package main

import (
	"fmt"
)

func main() {
	var a = "initial"
	fmt.Println(a)

	//========基本定义变量方式==========//
	/*定义变量的关键字 var
	b, c 变量名
	int 定义变量的类型
	= 1, 2 初始化变量
	*/
	var b, c int = 1, 2 //同时初始化多个变量,又叫平行赋值
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	var f = 1 //语言的自省
	fmt.Println(f)

	//========简约定义方式===========//
	/*
		:=   简短声明形式
		直接取代了var 和 type
	*/
	g := "test"
	fmt.Println(g)

	//==========INFO==============//
	/*
		:= 只能在函数内部，否则无法编译通过(而var没有限制)。
		所以 var 一般用来定义全局变量，:= 定义局部变量
	*/

}

/*
output:
$ go run variables.go
initial
1 2
true
0
1
test
*/
