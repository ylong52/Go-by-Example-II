/*
Go的基本语法结构
*/

//当前程序包名(一个可执行程序只有一个main包)
//一般建议package的名称和目录名保持一致
package main

//导入其他包
import "fmt"

//通过func关键词来进行函数的声明
//只有package名称为main的包才可以包含main函数
func main() {
	fmt.Println("Hello World")
}

/*============================
output:
$ go run hello-world.go
Hello World
*/
