//Go SHA1散列

package main

import (
	"crypto/sha1"
	"fmt"
)

/*
	SHA1散列常用于生成二进制文件或者文本块的短标识。
	Go在多个crypto/*包中实现了一系列散列函数
*/

func main() {
	s := "sha1 this string"

	//产生一个新的散列
	h := sha1.New()

	//写入要处理的字节，如果是字符串需要使用[]byte(s)来强制转换成字节数组
	h.Write([]byte(s))

	//Sum的参数可以用现有的字符串追加额外的字节切片，一般不需要
	bs := h.Sum(nil)

	fmt.Println(s)
	//SHA1经常以16进制输出，用%x来讲散列的结果格式化为16进制字符串
	fmt.Printf("%x\n", bs)
}

/*
output:
$ go run sha1-hashes.go
sha1 this string
cf23df2207d99a74fbe169e3eba035e633b65d94
*/
