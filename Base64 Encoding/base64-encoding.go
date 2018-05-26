//Go Base64编解码

package main

import (
	"encoding/base64"
	"fmt"
)

/*
	Go的内置包 encoding/base64 提供了base64编解码

	Go同时支持标准的和URL兼容的base64格式。
	编码需要使用[]byte类型的参数，所以要将字符串转换成此类型
*/

func main() {

	data := "abc123!?$*&()'-=@~"

	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)

	uDec, _ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}

/*
output:
$ go run base64-encoding.go
YWJjMTIzIT8kKiYoKSctPUB+
abc123!?$*&()'-=@~

YWJjMTIzIT8kKiYoKSctPUB-
abc123!?$*&()'-=@~
*/

/*
	标准base64编码和URL兼容base64编码的编码字符串存在稍许不同(后缀为+和-，但是两者都可以正确解码为原始字符串)
*/
