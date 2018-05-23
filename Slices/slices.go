// Go 的切片(数组切片)

package main

import (
	"fmt"
)

/*
	slice介绍
	数组的长度在定义之后无法再次修改；数组是值类型，每次传递都将产生一份副本。
	显然这种数据结构无法完全满足开发者的真实需求。
	在初始定义数组时，我们并不知道需要多大的数组，因此我们就需要“动态数组”。
	在Go里面这种数据结构叫slice，slice并不是真正意义上的动态数组，而是一个引用类型。
	slice总是指向一个底层array，slice的声明也可以像array一样，只是不需要长度，它是可变长的，可以随时往slice里面加数据。

	数组切片的数据结构
	1. 一个指向原生数组的指针(point): 指向数组中slice指定的开始位置
	2. 数组切片中的元素个数(len): 即slice的长度
	3. 数组切片已分配的存储空间(cap): 也就是slice开始位置到数组的最后位置的长度
*/

/*
	slice声明
	var myslice []int (声明时方括号没有任何数据)

	mySlice := {'a', 'b', 'c'} //直接创建并初始化包含三个元素的数组切片

	利用Go的内置函数 make 初始化切片
	aSlice := make([]int, 5) //创建一个初始元素个数为5的整型数值切片，元素初始值为0
	bSlice := make([]int, 5, 10) //创建一个初始元素个数为5的整型数组切片，元素初始值为0，并预留10个元素的存储空间
*/

func main() {

	//定义一个切片 元素类型为 string
	var aSlice []string
	//go自动初始化为nil，长度为0，地址为0， nil
	fmt.Printf("length:%v \t addr:%p \t isnil:%v", len(aSlice), aSlice, aSlice == nil)

	bSlice := []int{1, 2, 3, 4, 5}

	//slice元素遍历(range)
	/*range的返回值: 索引，元素的值*/
	for _, v := range bSlice {
		fmt.Println(v)
		break
	}

	//slice增加元素(append)
	bSlice = append(bSlice, 6, 7)
	fmt.Println(bSlice)
	insertSlice := make([]int, 5)
	bSlice = append(bSlice, insertSlice...)
	fmt.Println(bSlice)

	//slice删除指定索引元素
	var dSlice []int
	index := 1
	dSlice = append(bSlice[:index], bSlice[index+1:]...)
	fmt.Println(dSlice)
}

/*
output:
$ go run slices.go
length:0         addr:0x0        isnil:true1
[1 2 3 4 5 6 7]
[1 2 3 4 5 6 7 0 0 0 0 0]
[1 3 4 5 6 7 0 0 0 0 0]
*/
