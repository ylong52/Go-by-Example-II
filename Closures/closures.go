//Go 闭包

package main

/*

 */

import (
	"fmt"
)

//普通调用函数
func p(i int) {
	fmt.Println(i)
}

func main() {

	fmt.Println("=============普通函数调用=================")
	a := []int{1, 2, 3}
	for _, i := range a {
		fmt.Println(i)
		//defer: 后进先出， 所以参数传入顺序为3->2->1
		defer p(i) //最后执行
	}

	fmt.Println("=====================闭包==================")
	b := []int{1, 2, 3}
	for _, j := range b {
		fmt.Println(j)
		defer func() {
			fmt.Println(j)
		}()
	}
}

/*
output:
$ go run closures.go
=============普通函数调用=================
1
2
3
=====================闭包==================
1
2
3
3
3
3
3
2
1
*/

/*
	闭包里的非传递参数外部变量值是传引用的，在上面的闭包函数里，j 就是外部非闭包函数自己的参数，
	所以是相当于引用了外部的变量，j的值执行到第三次是3，闭包是地址引用，所以打印了三次j地址指向的值，
	所以是3, 3, 3
*/
