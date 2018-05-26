//Go 数字解析

package main

import (
	"fmt"
	"strconv"
)

/*
	Go内置的strconv包提供了数字解析功能
*/

func main() {
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	//自动腿短字符串所表示的数字的进制
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	//Atoi是一个基础的额10进制整型数转换函数
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	//输入参数错误时解析函数会返回一个错误
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}

/*
output:
$ go run number-parsing.go
1.234
123
456
789
135
strconv.Atoi: parsing "wat": invalid syntax
*/
