//Go 可变长参数函数

package main

import (
	"fmt"
)

//定义一个函数，可以传入任意数量的整型参数
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	//定义一个切片数组
	nums := []int{1, 2, 3, 4}

	//如果需要传入的参数在一个切片中，则要把切片展开传入
	// 此时的传入写法为: func(slice...)
	sum(nums...)
}

/*
output:
$ go run variadic-functions.go
[1 2] 3
[1 2 3] 6
[1 2 3 4] 10
*/

/*
要注意的是，可变长参数是函数定义的最后一个参数，比如上面的...int
*/
