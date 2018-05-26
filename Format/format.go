//Go 格式化输出

package main

import (
	"fmt"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}

	fmt.Printf("%v\n", p)

	fmt.Printf("%+v\n", p)

	fmt.Printf("%#v\n", p)

	fmt.Printf("%t\n", p)

	fmt.Printf("%t\n", true)

	fmt.Printf("%d\n", 123)

	fmt.Printf("%b\n", 14)

	fmt.Printf("%c\n", 33)

	fmt.Printf("%x\n", 456)

	fmt.Printf("%f\n", 78.9)

	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	fmt.Printf("%s\n", "\"string\"")

	fmt.Printf("%q\n", "\"string\"")

	fmt.Printf("%x\n", "hex this")

	fmt.Printf("%p\n", &p)

	fmt.Printf("|%6d|%6d|\n", 12, 345)

	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	fmt.Printf("|%-6s|%-6s|\n", "FOO", "B")

	fmt.Printf("|%6s|%6s|\n", "FOO", "B")

	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)
}

/*
output:
$ go run format.go
{1 2}
{x:1 y:2}
main.point{x:1, y:2}
{%!t(int=1) %!t(int=2)}
true
123
1110
!
1c8
78.900000
1.234000e+08
1.234000E+08
"string"
"\"string\""
6865782074686973
0xc042060080
|    12|   345|
|  1.20|  3.45|
|1.20  |3.45  |
|FOO   |B     |
|   FOO|     B|
a string
*/
