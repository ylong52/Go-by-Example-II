//Go (内置)排序

package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "b", "a"}
	//字符串排序
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	//int型排序
	sort.Ints(ints)
	fmt.Println("Ints:", ints)

	//排序检测(bool)
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)
}

/*
output:
$ go run sorting.go
Strings: [a b c]
Ints: [2 4 7]
Sorted: true
*/
