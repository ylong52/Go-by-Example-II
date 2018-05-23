//Go的if/else

package main

import "fmt"

func main() {
	for i := 6; i <= 9; i++ {
		fmt.Println("==========")
		if i%2 == 0 {
			fmt.Printf("%d 是偶数\n", i)
		} else {
			fmt.Printf("%d 是奇数\n", i)
		}
	}
}

/*
output:
$ go run if-else.go
==========
6 是偶数
==========
7 是奇数
==========
8 是偶数
==========
9 是奇数
*/

/*
Printf 和Println 都是fmt包中的公共方法
Printf: 可以打印出字符串和变量；
Println: 只可以打印出格式化的字符串,可以输出字符串类型的变量,不可以输出整形变量和整形
*/
