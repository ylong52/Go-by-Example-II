//Go的map

package main

import (
	"fmt"
)

/*
go的数组和切片都是顺序性列表，而map则是无序的
Go的map与python的字典类似
*/

func main() {

	//创建一个键为string类型，值为string类型的map
	m1 := make(map[string]string)
	m1["name"] = "apple"
	m1["category"] = "fruit"
	fmt.Println("this map is:", m1)

	//创建一个键为string类型，值为int类型的map
	m2 := make(map[string]int)
	m2["k1"] = 7
	m2["k2"] = 13
	v1 := m2["k1"]
	fmt.Println("v1:", v1)
	delete(m2, "k2")
	fmt.Println("after delete map is:", m2)
	_, prs := m2["k2"]
	fmt.Println("prs:", prs)

	//创建一个键为string类型，值为int类型的map,并赋值
	n := map[string]int{"for": 1, "bar": 2}
	fmt.Println("the new map is:", n)
}

/*
output:
$ go run maps.go
this map is: map[name:apple category:fruit]
v1: 7
after delete map is: map[k1:7]
prs: false
the new map is: map[for:1 bar:2]
*/
