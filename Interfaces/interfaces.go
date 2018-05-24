//Go 接口

package main

/*
	接口:方法的集合，不能包含变量
	接口将类型能做什么和如何做分离开，使得相同接口的变量在不同时刻表现出不同行为，这就是多态的本质
	接口申明方式:
	type InterfaceNamer interface {
		Meathod1() return_type
		Meathod2() return_type
	}
*/

import (
	"fmt"
	"math"
)

//Geometry 定义一个关于几何图形的接口，里面实现了几何图形的求面积和周长的方法
type Geometry interface {
	area() float64
	perim() float64
}

//Rect 定义一个矩形结构体
type Rect struct {
	width, height float64
}

//Circle 定义一个圆结构体
type Circle struct {
	radius float64
}

//实现一个求矩形面积的方法
func (r Rect) area() float64 {
	return r.height * r.width
}

//实现一个求矩形周长的方法
func (r Rect) perim() float64 {
	return 2*r.width + 2*r.height
}

//实现一个求圆面积的方法
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

//实现一个求圆周长的方法
func (c Circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := &Rect{3, 4}
	c := &Circle{5}

	measure(r)
	measure(c)
}

/*
output:
$ go run interfaces.go
&{3 4}
12
14
&{5}
78.53981633974483
31.41592653589793
*/
