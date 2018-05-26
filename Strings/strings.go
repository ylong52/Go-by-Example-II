//Go 标准库之strings

package main

import (
	"fmt"
	"strings"
)

var p = fmt.Println

func main() {

	//字符串是否有交集，返回值bool
	p("contains: ", strings.Contains("test", "es"))
	//统计某个字符串个数
	p("count: ", strings.Count("test", "t"))
	//前缀判断 返回值bool
	p("hasprefix: ", strings.HasPrefix("test", "te"))
	//后缀判断 返回值bool
	p("hassuffix: ", strings.HasSuffix("test", "st"))
	//索引
	p("index: ", strings.Index("test", "e"))
	//连接数组，返回值是字符串
	p("join: ", strings.Join([]string{"A", "B"}, "-"))
	//重复某个字符
	p("repeat: ", strings.Repeat("A", 5))
	//字符串替换，str，old，new，n(int 限制)
	p("replace: ", strings.Replace("FOO", "O", "0", -1))
	p("replace: ", strings.Replace("FOO", "O", "0", 1))
	//字符串分割，返回值是一个数组
	p("split: ", strings.Split("A-B-C-D-E", "-"))
	//转小写
	p("tolower: ", strings.ToLower("TEST"))
	//转大小
	p("toupper: ", strings.ToUpper("test"))

	//字符串长度
	p("len: ", len("test"))
	//字符串切片返回值是unicode 字节码
	p("char: ", "test"[1])
}

/*
outut:
$ go run strings.go
contains:  true
count:  2
hasprefix:  true
hassuffix:  true
index:  1
join:  A-B
repeat:  AAAAA
replace:  F00
replace:  F0O
split:  [A B C D E]
tolower:  test
toupper:  TEST
len:  4
char:  101
*/
