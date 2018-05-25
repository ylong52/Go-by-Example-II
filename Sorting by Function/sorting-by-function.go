//Go 自定义排序

package main

import (
	"fmt"
	"sort"
)

/*
golang内置的排序包提供了自定义排序的接口
sort包在内部实现了四种基本的排序算法:插入排序、归并排序、堆排序、快速排序
sort包会依据实际数据自动选择最优的排序算法。

调用sort.Sort()来实现自定义排序，只需要我们的类型实现Interface接口类型中的三个方法即可

sort包内部对于[]int类型的排序方法
//首先定义一个[]int类型的别名IntSlice
type IntSlice []int
//获取此slice的长度
func (p IntSlice) Len() int {
	return len(p)
}
//比较两个元素大小 升序
func (p IntSlice) Less(i, j int) bool {
	return p[i] < p[j]
}
//交换数据
func (p IntSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
//sort.Ints()内部调用Sort()方法实现排序
//注意要先将[]int转换为IntSlice类型,因为此类型才实现了Interface的三个方法
func Ints(a []int) {
	Sort(IntSlice(a))
}
*/

type byLength []string

func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	//将fruits转换成byLength类型来调用排序接口
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}

/*
output:
$ go run sorting-by-function.go
[kiwi peach banana]
*/
