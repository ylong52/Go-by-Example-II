//Go for循环

package main

import (
	"fmt"
)

func main() {

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println("==========")

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	fmt.Println("==========")

	//循环一次则break跳出循环
	for {
		fmt.Println("loop")
		break
	}
	fmt.Println("==========")

	for n := 0; n <= 5; n++ {
		//如果n为2则跳到下一次循环
		if n == 2 {
			continue
		}
		fmt.Println(n)
	}
}

/*
output:
$ go run for.go
1
2
3
==========
7
8
9
==========
loop
==========
0
1
3
4
5

*/
