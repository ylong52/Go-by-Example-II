//Go 读文件

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	//文件不存在会报错，此时可以看到ioutil是已协程的方式进行的
	dat, err := ioutil.ReadFile("./dat")
	check(err)
	fmt.Print(string(dat))

	f, err := os.Open("./dat")
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))
}

/*
output:
$ go run reading-files.go
test4 bytes: test
*/
