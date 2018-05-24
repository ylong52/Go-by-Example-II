//Go 方法method

package main

/*
	通过显示说明receiver来实现与某个类型的结合
	只能为同一个包中的类型定义方法
	receiver可以是类型的值或指针
	不存在方法重载
	可以使用值或指针来调用方法，编译器会自动完成转换
	从某种意义上说，方法是函数的语法糖，因为receiver其实就是方法所接收的第一个参数
	如果外部结构和嵌入结构存在同名方法，则优先调用外部结构方法
	类型别名不会拥有底层类型附带的方法
	方法可以调用结构中的非公开字段
*/

import (
	"fmt"
)

//Rect 定义一个结构体
type Rect struct {
	width, height int
}

//receiver:指针传递
func (r *Rect) area() int {
	return r.width * r.height
}

//receiver:值传递
func (r Rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := &Rect{10, 5}

	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())
}

/*
output:
$ go run methods.go
area: 50
perim: 30
*/
