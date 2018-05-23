//Go中的range(遍历)

package main

import (
	"fmt"
)

func main() {

	//创建一个数组
	nums := []int{2, 3, 4}
	//定义一个变量
	sum := 0
	//range的返回值为索引和数组中的元素
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	//创建一个字典
	//rang在map中的使用，索引为键，元素为值
	kvs := map[string]string{"a": "apple", "b": "banana"}
	//遍历键值对
	for k, v := range kvs {
		fmt.Printf("%s:%s", k, v)
	}
	//遍历键
	for k := range kvs {
		fmt.Println("key:", k)
	}

	//遍历字符串,返回Unicode代码点
	// 第一个返回值是每个字符的起始字节的索引，第二个是字符代码点，
	// 因为Go的字符串是由字节组成的，多个字节组成一个rune类型字符。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

/*
output:
$ go run range.go
sum: 9
index: 1
a:appleb:bananakey: a
key: b
0 103
1 111
*/
