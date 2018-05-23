//Go的值的类型，str,int,floats,bool,etc

package main

import "fmt"

func main() {
	fmt.Println("go" + "lang")

	fmt.Println("1+1=", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

/*
output:
$ go run values.go
golang
1+1= 2
7.0/3.0 = 2.3333333333333335
false
true
false
*/

/*
golang的浮点数不精准
出现浮点数不精确的原因是，浮点数储存至内存中时，
2的-1、-2……-n次方不能精确的表示小数部分，
所以再把这个数从地址中取出来进行计算就出现了偏差
*/
