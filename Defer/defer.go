//Go 异常处理之defer

package main

/*
	defer用来添加函数结束时执行的语句，是动态的添加。
*/

import (
	"fmt"
	"os"
)

func main() {

	f := CreateFile("./defer.txt")
	defer CloseFile(f)
	WriteFile(f)
}

//CreateFile 新建文件
func CreateFile(p string) *os.File {
	fmt.Println("Creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

//CloseFile 关闭文件
func CloseFile(f *os.File) {
	fmt.Println("Closing")
	f.Close()
}

//WriteFile 写文件
func WriteFile(f *os.File) {
	fmt.Println("Writing")
	fmt.Fprintln(f, "Data")
}

/*
output:
$ go run defer.go
Creating
Writing
Closing
*/
