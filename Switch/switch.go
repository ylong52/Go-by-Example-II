//Go 的选择语句(switch)
//根据传入的条件不同，选择语句会执行不同的语句

package main

import (
	"fmt"
)

var i int

func main() {

	fmt.Println("输入一个数字：")
	//获取控制台输入
	fmt.Scanln(&i)

	//默认执行第一个case
	switch i {
	case 0:
		fmt.Println("你输入的是零")
	case 1:
		fmt.Println("你输入的是一")
	case 2:
		//跳到下一个case
		fallthrough
	case 3, 4, 5:
		fmt.Println("你输入的是三或者四或者五")
	//没有满足的case则执行default
	default:
		fmt.Println("你输入的数字大于5")
	}

}

/*
output:
$ go run switch.go
输入一个数字：
你输入的是零

$ go run switch.go
输入一个数字：
1
你输入的是一

$ go run switch.go
输入一个数字：
2
你输入的是三或者四或者五

$ go run switch.go
输入一个数字：
3
你输入的是三或者四或者五

$ go run switch.go
输入一个数字：
6
你输入的数字大于5
*/

/*
在使用switch结构时，我们需要注意以下几点:
1. 左花括号必须与switch同行
2. 条件表达式不限制为常量或者整数
3. 单个case中，可以出现多个结果选项
4. 与C语言不同的是，Go语言不需要break来明确退出一个case
5. 只有在case中明确添加fallthrough关键字，才会执行紧跟的下一个case
6. 可以不设定switch之后的表达式，在这种情况下整个switch结构个多个if/else 的逻辑作用等同
*/
